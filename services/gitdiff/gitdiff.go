	"code.gitea.io/gitea/models/db"
	user_model "code.gitea.io/gitea/models/user"
	"code.gitea.io/gitea/modules/analyze"
	file     *DiffFile
	language := ""
	if diffSection.file != nil {
		language = diffSection.file.Language
	}

			return template.HTML(highlight.Code(diffSection.FileName, language, diffLine.Content[1:]))
			return template.HTML(highlight.Code(diffSection.FileName, language, diffLine.Content[1:]))
			return template.HTML(highlight.Code(diffSection.FileName, language, diffLine.Content[1:]))
		return template.HTML(highlight.Code(diffSection.FileName, language, diffLine.Content))
	diffRecord := diffMatchPatch.DiffMain(highlight.Code(diffSection.FileName, language, diff1[1:]), highlight.Code(diffSection.FileName, language, diff2[1:]), true)
	IsGenerated             bool
	IsVendored              bool
	Language                string
	Start, End                             string
func (diff *Diff) LoadComments(issue *models.Issue, currentUser *user_model.User) error {
func ParsePatch(maxLines, maxLineCharacters, maxFiles int, reader io.Reader, skipToFile string) (*Diff, error) {
	log.Debug("ParsePatch(%d, %d, %d, ..., %s)", maxLines, maxLineCharacters, maxFiles, skipToFile)
	skipping := skipToFile != ""

			return diff, fmt.Errorf("invalid first file line: %s", line)
		if maxFiles > -1 && len(diff.Files) >= maxFiles {
			lastFile := createDiffFile(diff, line)
			diff.End = lastFile.Name
			_, err := io.Copy(io.Discard, reader)
				return diff, fmt.Errorf("error during io.Copy: %w", err)
		if skipping {
			if curFile.Name != skipToFile {
				line, err = skipToNextDiffHead(input)
				if err != nil {
					if err == io.EOF {
						return diff, nil
					}
					return diff, err
				}
				continue
			}
			skipping = false
		}

						return diff, fmt.Errorf("unable to ReadLine: %w", err)
func skipToNextDiffHead(input *bufio.Reader) (line string, err error) {
	// need to skip until the next cmdDiffHead
	isFragment, wasFragment := false, false
	var lineBytes []byte
	for {
		lineBytes, isFragment, err = input.ReadLine()
		if err != nil {
			return
		}
		if wasFragment {
			wasFragment = isFragment
			continue
		}
		if bytes.HasPrefix(lineBytes, []byte(cmdDiffHead)) {
			break
		}
		wasFragment = isFragment
	}
	line = string(lineBytes)
	if isFragment {
		var tail string
		tail, err = input.ReadString('\n')
		if err != nil {
			return
		}
		line += tail
	}
	return
}

				err = fmt.Errorf("unable to ReadLine: %w", err)
			err = fmt.Errorf("unable to ReadLine: %w", err)
			if maxLines > -1 && curFileLinesCount >= maxLines {
					err = fmt.Errorf("unable to ReadLine: %w", err)
			curSection = &DiffSection{file: curFile}
			if maxLines > -1 && curFileLinesCount >= maxLines {
				err = fmt.Errorf("unexpected line in hunk: %s", string(lineBytes))
			if maxLines > -1 && curFileLinesCount >= maxLines {
				curSection = &DiffSection{file: curFile}
			if maxLines > -1 && curFileLinesCount >= maxLines {
				curSection = &DiffSection{file: curFile}
			if maxLines > -1 && curFileLinesCount >= maxLines {
				curSection = &DiffSection{file: curFile}
			err = fmt.Errorf("unexpected line in hunk: %s", string(lineBytes))
					err = fmt.Errorf("unable to ReadLine: %w", err)
				count, err := db.Count(m)

// DiffOptions represents the options for a DiffRange
type DiffOptions struct {
	BeforeCommitID     string
	AfterCommitID      string
	SkipTo             string
	MaxLines           int
	MaxLineCharacters  int
	MaxFiles           int
	WhitespaceBehavior string
	DirectComparison   bool
}

// GetDiff builds a Diff between two commits of a repository.
func GetDiff(gitRepo *git.Repository, opts *DiffOptions, files ...string) (*Diff, error) {
	commit, err := gitRepo.GetCommit(opts.AfterCommitID)
	timeout := time.Duration(setting.Git.Timeout.Default) * time.Second
	ctx, _, finished := process.GetManager().AddContextTimeout(gitRepo.Ctx, timeout, fmt.Sprintf("GetDiffRange [repo_path: %s]", repoPath))
	defer finished()

	argsLength := 6
	if len(opts.WhitespaceBehavior) > 0 {
		argsLength++
	}
	if len(opts.SkipTo) > 0 {
		argsLength++
	}
	if len(files) > 0 {
		argsLength += len(files) + 1
	}
	diffArgs := make([]string, 0, argsLength)
	if (len(opts.BeforeCommitID) == 0 || opts.BeforeCommitID == git.EmptySHA) && commit.ParentCount() == 0 {
		diffArgs = append(diffArgs, "diff", "--src-prefix=\\a/", "--dst-prefix=\\b/", "-M")
		if len(opts.WhitespaceBehavior) != 0 {
			diffArgs = append(diffArgs, opts.WhitespaceBehavior)
		diffArgs = append(diffArgs, opts.AfterCommitID)
		actualBeforeCommitID := opts.BeforeCommitID
		diffArgs = append(diffArgs, "diff", "--src-prefix=\\a/", "--dst-prefix=\\b/", "-M")
		if len(opts.WhitespaceBehavior) != 0 {
			diffArgs = append(diffArgs, opts.WhitespaceBehavior)
		diffArgs = append(diffArgs, opts.AfterCommitID)
		opts.BeforeCommitID = actualBeforeCommitID
	}

	// In git 2.31, git diff learned --skip-to which we can use to shortcut skip to file
	// so if we are using at least this version of git we don't have to tell ParsePatch to do
	// the skipping for us
	parsePatchSkipToFile := opts.SkipTo
	if opts.SkipTo != "" && git.CheckGitVersionAtLeast("2.31") == nil {
		diffArgs = append(diffArgs, "--skip-to="+opts.SkipTo)
		parsePatchSkipToFile = ""

	if len(files) > 0 {
		diffArgs = append(diffArgs, "--")
		diffArgs = append(diffArgs, files...)
	}

	cmd := exec.CommandContext(ctx, git.GitExecutable, diffArgs...)

		return nil, fmt.Errorf("error creating StdoutPipe: %w", err)
		return nil, fmt.Errorf("error during Start: %w", err)
	diff, err := ParsePatch(opts.MaxLines, opts.MaxLineCharacters, opts.MaxFiles, stdout, parsePatchSkipToFile)
		return nil, fmt.Errorf("unable to ParsePatch: %w", err)
	}
	diff.Start = opts.SkipTo

	var checker *git.CheckAttributeReader

	if git.CheckGitVersionAtLeast("1.7.8") == nil {
		indexFilename, worktree, deleteTemporaryFile, err := gitRepo.ReadTreeToTemporaryIndex(opts.AfterCommitID)
		if err == nil {
			defer deleteTemporaryFile()

			checker = &git.CheckAttributeReader{
				Attributes: []string{"linguist-vendored", "linguist-generated", "linguist-language", "gitlab-language"},
				Repo:       gitRepo,
				IndexFile:  indexFilename,
				WorkTree:   worktree,
			}
			ctx, cancel := context.WithCancel(git.DefaultContext)
			if err := checker.Init(ctx); err != nil {
				log.Error("Unable to open checker for %s. Error: %v", opts.AfterCommitID, err)
			} else {
				go func() {
					err := checker.Run()
					if err != nil && err != ctx.Err() {
						log.Error("Unable to open checker for %s. Error: %v", opts.AfterCommitID, err)
					}
					cancel()
				}()
			}
			defer func() {
				cancel()
			}()
		}


		gotVendor := false
		gotGenerated := false
		if checker != nil {
			attrs, err := checker.CheckPath(diffFile.Name)
			if err == nil {
				if vendored, has := attrs["linguist-vendored"]; has {
					if vendored == "set" || vendored == "true" {
						diffFile.IsVendored = true
						gotVendor = true
					} else {
						gotVendor = vendored == "false"
					}
				}
				if generated, has := attrs["linguist-generated"]; has {
					if generated == "set" || generated == "true" {
						diffFile.IsGenerated = true
						gotGenerated = true
					} else {
						gotGenerated = generated == "false"
					}
				}
				if language, has := attrs["linguist-language"]; has && language != "unspecified" && language != "" {
					diffFile.Language = language
				} else if language, has := attrs["gitlab-language"]; has && language != "unspecified" && language != "" {
					diffFile.Language = language
				}
			} else {
				log.Error("Unexpected error: %v", err)
			}
		}

		if !gotVendor {
			diffFile.IsVendored = analyze.IsVendor(diffFile.Name)
		}
		if !gotGenerated {
			diffFile.IsGenerated = analyze.IsGenerated(diffFile.Name)
		}

		tailSection := diffFile.GetTailSection(gitRepo, opts.BeforeCommitID, opts.AfterCommitID)
		return nil, fmt.Errorf("error during cmd.Wait: %w", err)
	}

	separator := "..."
	if opts.DirectComparison {
		separator = ".."
	shortstatArgs := []string{opts.BeforeCommitID + separator + opts.AfterCommitID}
	if len(opts.BeforeCommitID) == 0 || opts.BeforeCommitID == git.EmptySHA {
		shortstatArgs = []string{git.EmptyTreeSHA, opts.AfterCommitID}
		shortstatArgs = []string{opts.BeforeCommitID, opts.AfterCommitID}
		setting.Git.MaxGitDiffLineCharacters, setting.Git.MaxGitDiffFiles, strings.NewReader(c.Patch), "")