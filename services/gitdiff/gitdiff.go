	"io/ioutil"
	"os/exec"
	"code.gitea.io/gitea/modules/process"
func getLineContent(content string) string {
		return html.EscapeString(content)
	return "<br>"
var unfinishedtagRegex = regexp.MustCompile(`<[^>]*$`)
var trailingSpanRegex = regexp.MustCompile(`<span\s*[[:alpha:]="]*?[>]?$`)
var entityRegex = regexp.MustCompile(`&[#]*?[0-9[:alpha:]]*$`)

func diffToHTML(fileName string, diffs []diffmatchpatch.Diff, lineType DiffLineType) template.HTML {
	return template.HTML(buf.Bytes())
func (diffSection *DiffSection) GetComputedInlineDiffFor(diffLine *DiffLine) template.HTML {
		return template.HTML(getLineContent(diffLine.Content[1:]))
		return template.HTML(getLineContent(diffLine.Content[1:]))
			return template.HTML(highlight.Code(diffSection.FileName, diffLine.Content[1:]))
			return template.HTML(highlight.Code(diffSection.FileName, diffLine.Content[1:]))
			return template.HTML(highlight.Code(diffSection.FileName, diffLine.Content[1:]))
		return template.HTML(highlight.Code(diffSection.FileName, diffLine.Content))
	diffRecord := diffMatchPatch.DiffMain(highlight.Code(diffSection.FileName, diff1[1:]), highlight.Code(diffSection.FileName, diff2[1:]), true)
		}}

func (diff *Diff) LoadComments(issue *models.Issue, currentUser *models.User) error {
	allComments, err := models.FetchCodeComments(issue, currentUser)
func ParsePatch(maxLines, maxLineCharacters, maxFiles int, reader io.Reader) (*Diff, error) {
			return diff, fmt.Errorf("Invalid first file line: %s", line)
		// TODO: Handle skipping first n files
		if len(diff.Files) >= maxFiles {
			_, err := io.Copy(ioutil.Discard, reader)
				return diff, fmt.Errorf("Copy: %v", err)
					// "--- a\t\n" without the qoutes.
						return diff, fmt.Errorf("Unable to ReadLine: %v", err)
	var diffLineTypeBuffers = make(map[DiffLineType]*bytes.Buffer, 3)
	var diffLineTypeDecoders = make(map[DiffLineType]*encoding.Decoder, 3)
				err = fmt.Errorf("Unable to ReadLine: %v", err)
			err = fmt.Errorf("Unable to ReadLine: %v", err)
			if curFileLinesCount >= maxLines {
					err = fmt.Errorf("Unable to ReadLine: %v", err)
			curSection = &DiffSection{}
			if curFileLinesCount >= maxLines {
				err = fmt.Errorf("Unexpected line in hunk: %s", string(lineBytes))
			if curFileLinesCount >= maxLines {
				curSection = &DiffSection{}
			if curFileLinesCount >= maxLines {
				curSection = &DiffSection{}
			if curFileLinesCount >= maxLines {
				curSection = &DiffSection{}
			err = fmt.Errorf("Unexpected line in hunk: %s", string(lineBytes))
					err = fmt.Errorf("Unable to ReadLine: %v", err)
				count, err := models.Count(m)
// GetDiffRangeWithWhitespaceBehavior builds a Diff between two commits of a repository.
func GetDiffRangeWithWhitespaceBehavior(gitRepo *git.Repository, beforeCommitID, afterCommitID string, maxLines, maxLineCharacters, maxFiles int, whitespaceBehavior string) (*Diff, error) {
	commit, err := gitRepo.GetCommit(afterCommitID)
	ctx, cancel := context.WithTimeout(git.DefaultContext, time.Duration(setting.Git.Timeout.Default)*time.Second)
	defer cancel()
	var cmd *exec.Cmd
	if (len(beforeCommitID) == 0 || beforeCommitID == git.EmptySHA) && commit.ParentCount() == 0 {
		diffArgs := []string{"diff", "--src-prefix=\\a/", "--dst-prefix=\\b/", "-M"}
		if len(whitespaceBehavior) != 0 {
			diffArgs = append(diffArgs, whitespaceBehavior)
		diffArgs = append(diffArgs, afterCommitID)
		cmd = exec.CommandContext(ctx, git.GitExecutable, diffArgs...)
		actualBeforeCommitID := beforeCommitID
		diffArgs := []string{"diff", "--src-prefix=\\a/", "--dst-prefix=\\b/", "-M"}
		if len(whitespaceBehavior) != 0 {
			diffArgs = append(diffArgs, whitespaceBehavior)
		diffArgs = append(diffArgs, afterCommitID)
		cmd = exec.CommandContext(ctx, git.GitExecutable, diffArgs...)
		beforeCommitID = actualBeforeCommitID
	cmd.Dir = repoPath
	cmd.Stderr = os.Stderr
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, fmt.Errorf("StdoutPipe: %v", err)
	if err = cmd.Start(); err != nil {
		return nil, fmt.Errorf("Start: %v", err)
	pid := process.GetManager().Add(fmt.Sprintf("GetDiffRange [repo_path: %s]", repoPath), cancel)
	defer process.GetManager().Remove(pid)
	diff, err := ParsePatch(maxLines, maxLineCharacters, maxFiles, stdout)
		return nil, fmt.Errorf("ParsePatch: %v", err)
		tailSection := diffFile.GetTailSection(gitRepo, beforeCommitID, afterCommitID)
	if err = cmd.Wait(); err != nil {
		return nil, fmt.Errorf("Wait: %v", err)
	shortstatArgs := []string{beforeCommitID + "..." + afterCommitID}
	if len(beforeCommitID) == 0 || beforeCommitID == git.EmptySHA {
		shortstatArgs = []string{git.EmptyTreeSHA, afterCommitID}
	diff.NumFiles, diff.TotalAddition, diff.TotalDeletion, err = git.GetDiffShortStat(repoPath, shortstatArgs...)
		shortstatArgs = []string{beforeCommitID, afterCommitID}
		diff.NumFiles, diff.TotalAddition, diff.TotalDeletion, err = git.GetDiffShortStat(repoPath, shortstatArgs...)
// GetDiffCommitWithWhitespaceBehavior builds a Diff representing the given commitID.
// The whitespaceBehavior is either an empty string or a git flag
func GetDiffCommitWithWhitespaceBehavior(gitRepo *git.Repository, commitID string, maxLines, maxLineCharacters, maxFiles int, whitespaceBehavior string) (*Diff, error) {
	return GetDiffRangeWithWhitespaceBehavior(gitRepo, "", commitID, maxLines, maxLineCharacters, maxFiles, whitespaceBehavior)
}

		setting.Git.MaxGitDiffLineCharacters, setting.Git.MaxGitDiffFiles, strings.NewReader(c.Patch))
		"":              ""}