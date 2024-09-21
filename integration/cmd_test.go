package integration

import ("testing"
"filedownloader/shared/tests"
"filedownloader/cmd")

func Test_Run_ShouldDownloadFileFromServerAndSaveLocally(t *testing.T) {
	mockWebServer := tests.StartMockWebServer("sample response data")
	defer mockWebServer.Close()

	testDownloadUrl := mockWebServer.URL + "/download/file"
	tests.SetMockCliInput(testDownloadUrl)

	cmd.Run()

	
}