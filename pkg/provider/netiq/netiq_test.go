package netiq

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"github.com/stretchr/testify/require"
	"github.com/versent/saml2aws/v2/pkg/page"
	"io/ioutil"
	"net/url"
	"testing"
)

func TestIsSAMLResponsePositive(t *testing.T) {
	//given
	samlResponseData, err := ioutil.ReadFile("responses/samlRespose.html")
	require.Nil(t, err)

	//when
	samlRespDoc, err := goquery.NewDocumentFromReader(bytes.NewReader(samlResponseData))
	require.Nil(t, err)

	//then
	require.True(t, isSAMLResponse(samlRespDoc))
}

func TestIsSAMLResponseNegative(t *testing.T) {
	//given
	getToContentData, err := ioutil.ReadFile("responses/getToContent.html")
	require.Nil(t, err)

	//when
	getToContentDoc, err := goquery.NewDocumentFromReader(bytes.NewReader(getToContentData))
	require.Nil(t, err)

	//then
	require.False(t, isSAMLResponse(getToContentDoc))
}

func TestExtractSAMLAssertion(t *testing.T) {
	//given
	samlResponseData, err := ioutil.ReadFile("responses/samlRespose.html")
	require.Nil(t, err)
	expectedResult := "PHNhbWxwOlJlc3BvbnNlIHhtbG5zOnNhbWxwPSJ1cm46b2FzaXM6bmFtZXM6dGM6U0FNTDoyLjA6cHJvdG9jb2wiIHhtbG5zOnNhbWw9InVybjpvYXNpczpuYW1lczp0YzpTQU1MOjIuMDphc3NlcnRpb24iIERlc3RpbmF0aW9uPSJodHRwczovL3NpZ25pbi5hd3MuYW1hem9uLmNvbS9zYW1sIiBJRD0iaWRteFVmZUR5dVJhODlaTWtEOUg3Z3pKQl9tUTQiIElzc3VlSW5zdGFudD0iMjAyMC0wMy0wNVQwMjowNToxOFoiIFZlcnNpb249IjIuMCI+PHNhbWw6SXNzdWVyPmh0dHBzOi8vbG9naW4uYXV0aGJyaWRnZS53ZXN0cGFjZ3JvdXAuY29tL25pZHAvc2FtbDIvbWV0YWRhdGE8L3NhbWw6SXNzdWVyPjxzYW1scDpTdGF0dXM+PHNhbWxwOlN0YXR1c0NvZGUgVmFsdWU9InVybjpvYXNpczpuYW1lczp0YzpTQU1MOjIuMDpzdGF0dXM6U3VjY2VzcyIvPjwvc2FtbHA6U3RhdHVzPjxzYW1sOkFzc2VydGlvbiBJRD0iaWRpWmdNQTlZUlBGYURBSjdWaGZnTHpsckFLWjQiIElzc3VlSW5zdGFudD0iMjAyMC0wMy0wNVQwMjowNToxOFoiIFZlcnNpb249IjIuMCI+PHNhbWw6SXNzdWVyPmh0dHBzOi8vbG9naW4uYXV0aGJyaWRnZS53ZXN0cGFjZ3JvdXAuY29tL25pZHAvc2FtbDIvbWV0YWRhdGE8L3NhbWw6SXNzdWVyPjxkczpTaWduYXR1cmUgeG1sbnM6ZHM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvMDkveG1sZHNpZyMiPjxkczpTaWduZWRJbmZvPjxDYW5vbmljYWxpemF0aW9uTWV0aG9kIHhtbG5zPSJodHRwOi8vd3d3LnczLm9yZy8yMDAwLzA5L3htbGRzaWcjIiBBbGdvcml0aG09Imh0dHA6Ly93d3cudzMub3JnLzIwMDEvMTAveG1sLWV4Yy1jMTRuIyIvPjxkczpTaWduYXR1cmVNZXRob2QgQWxnb3JpdGhtPSJodHRwOi8vd3d3LnczLm9yZy8yMDAxLzA0L3htbGRzaWctbW9yZSNyc2Etc2hhMjU2Ii8+PGRzOlJlZmVyZW5jZSBVUkk9IiNpZGlaZ01BOVlSUEZhREFKN1ZoZmdMemxyQUtaNCI+PGRzOlRyYW5zZm9ybXM+PGRzOlRyYW5zZm9ybSBBbGdvcml0aG09Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvMDkveG1sZHNpZyNlbnZlbG9wZWQtc2lnbmF0dXJlIi8+PGRzOlRyYW5zZm9ybSBBbGdvcml0aG09Imh0dHA6Ly93d3cudzMub3JnLzIwMDEvMTAveG1sLWV4Yy1jMTRuIyIvPjwvZHM6VHJhbnNmb3Jtcz48ZHM6RGlnZXN0TWV0aG9kIEFsZ29yaXRobT0iaHR0cDovL3d3dy53My5vcmcvMjAwMS8wNC94bWxlbmMjc2hhMjU2Ii8+PERpZ2VzdFZhbHVlIHhtbG5zPSJodHRwOi8vd3d3LnczLm9yZy8yMDAwLzA5L3htbGRzaWcjIj45aDNZNUdxQlpvRFZ4K3czU2t0VmhKYzd5YXF3SU5FSFljRVlZWCt2RVMwPTwvRGlnZXN0VmFsdWU+PC9kczpSZWZlcmVuY2U+PC9kczpTaWduZWRJbmZvPjxTaWduYXR1cmVWYWx1ZSB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC8wOS94bWxkc2lnIyI+CmhZcUxFWVU3Wmc2eUFHZVRwVzNUSGJDTzlrNTRNaU1iRC95NmZ5ZW9SYThtbUxQYnpxaUtYR1NlQUMxT0t5UnRvanNLQStCbVpCTzQKQkl1bDRtajM1KzY4d3luZU9Gc3pSOVYyMlphSWJnenh2Ny8wR1FNL0s2ZEJpakdGc1NxcXZONUlBTi9DZDhJUDU3N1pyM2dNQ1ZqcApnQ0k4dTlQUktOclJHdmhRUmc4OW42bXdLcml4UnZkVHh5WXArT3R1VGh4NHkrcytkOGYyNEhxOVZ0M1AwckljSWJNelplMitWVmkzCmV5NVNDbnQzK3I1QVE4MHdhSHlPU3lld21ORHJNWHhrN3pWdW9XeVFrbFR3Y1ZQQUdZTDU1UXgrZG1mZ1FCU1JOZzN1VzM3T3czZkEKekZIc1I2eUhLdkc4dFRMWHVoVkp6am85OXV2MlF4Zi9nNFc3MGc9PQo8L1NpZ25hdHVyZVZhbHVlPjxkczpLZXlJbmZvPjxkczpYNTA5RGF0YT48ZHM6WDUwOUNlcnRpZmljYXRlPgpNSUlJcHpDQ0I0K2dBd0lCQWdJVEh3QUF2dWs4YUd1UzVzRkdYQUFBQUFDKzZUQU5CZ2txaGtpRzl3MEJBUXNGQURDQnBqRUxNQWtHCkExVUVCaE1DUVZVeEREQUtCZ05WQkFnVEEwNVRWekVQTUEwR0ExVUVCeE1HVTNsa2JtVjVNU1F3SWdZRFZRUUtFeHRYWlhOMGNHRmoKSUVKaGJtdHBibWNnUTI5eWNHOXlZWFJwYjI0eEx6QXRCZ05WQkFzVEprUnBaMmwwWVd3Z1EyVnlkR2xtYVdOaGRHVnpJRk5sWTNWeQphWFI1SUZObGNuWnBZMlZ6TVNFd0h3WURWUVFERXhoWFpYTjBjR0ZqSUZOSVFUSWdVMU5NSUVOQklGZFRSRU13SGhjTk1qQXdNakl3Ck1qQXlOVFU1V2hjTk1qSXdNakU1TWpBeU5UVTVXakNCdERFTE1Ba0dBMVVFQmhNQ1FWVXhHREFXQmdOVkJBZ1REMDVsZHlCVGIzVjAKYUNCWFlXeGxjekVQTUEwR0ExVUVCeE1HVTNsa2JtVjVNU1F3SWdZRFZRUUtFeHRYWlhOMGNHRmpJRUpoYm10cGJtY2dRMjl5Y0c5eQpZWFJwYjI0eEhUQWJCZ05WQkFzVEZGZENReUJKVTFOUUlFbEJUU0JRY205bmNtRnRNVFV3TXdZRFZRUURFeXh1WVcwdGMzTndMVUZYClV5MWhkWFJvWW5KcFpHZGxMWEJ5YjJRdWQyVnpkSEJoWTJkeWIzVndMbU52YlRDQ0FTSXdEUVlKS29aSWh2Y05BUUVCQlFBRGdnRVAKQURDQ0FRb0NnZ0VCQUpQMWtKQW81THgzWGRFRkIrcGxidE16dy9McUU3UnUzMnFXdkFVMUJ3YUVTQ2ZQak5KVW9Ga0RrS3NNeUE0WQpDcjYrcWhwcEQ0K0tlZC9tMjhmVTdmbWFVNm1hWUlsbFpsWnBnTFZqV3BNcDNQdVF1UnFtL0w4dlcvZHpURkdoOTZxb0pNR1paWWZ5CnYwaVVlM2wrdmpuNXFORjNBdERleDkrSWdHZUx3OHZ5L1NRRmFWeUlOOTNwYW4rM0hERDM0WFlTZmVidFBoQXppbGM4c3BtUmI3S2MKS2VsYmhDQjF4RVRaM29nODZCYzRJTzI5YXFVbXhZN01ybTVmQzdBcFpHYXVXWG55QWZEaGJkRlF3NGQ2elFOQWZpdGNKSk1CbW8rTgpJS1VpMzBBc0lCbHpxaVNNTTdNTFU3U3ZhcFN3QTdOTXNEWlhKR1FIMklqNVcyVWFuTDBDQXdFQUFhT0NCTHd3Z2dTNE1CMEdBMVVkCkRnUVdCQlJiblI4TnA3OHZKS05KY29ETWRVUGw1aDdYc3pBZkJnTlZIU01FR0RBV2dCUmRmdDhhVnM0RGhONFJiS1Zic3QrZUtOeUEKWmpDQ0FXY0dBMVVkSHdTQ0FWNHdnZ0ZhTUlJQlZxQ0NBVktnZ2dGT2htNW9kSFJ3T2k4dmQySmpZMkV1Y0d0cE1pNXpjbll1ZDJWegpkSEJoWXk1amIyMHVZWFV2UTBSUUwxZGxjM1J3WVdNbE1qQlRTRUV5SlRJd1UxTk1KVEl3UTBFbE1qQlhVMFJETDFkbGMzUndZV01sCk1qQlRTRUV5SlRJd1UxTk1KVEl3UTBFbE1qQlhVMFJETG1OeWJJYUIyMnhrWVhBNkx5OHZRMDQ5VjJWemRIQmhZeVV5TUZOSVFUSWwKTWpCVFUwd2xNakJEUVNVeU1GZFRSRU1zUTA0OVlYVXlNREEwYzNBeE1qWTFMRU5PUFVORVVDeERUajFRZFdKc2FXTWxNakJMWlhrbApNakJUWlhKMmFXTmxjeXhEVGoxVFpYSjJhV05sY3l4RFRqMURiMjVtYVdkMWNtRjBhVzl1TEVSRFBYZGlZMkYxTEVSRFBYZGxjM1J3CllXTXNSRU05WTI5dExFUkRQV0YxUDJObGNuUnBabWxqWVhSbFVtVjJiMk5oZEdsdmJreHBjM1EvWW1GelpUOXZZbXBsWTNSRGJHRnoKY3oxalVreEVhWE4wY21saWRYUnBiMjVRYjJsdWREQ0NBVjhHQ0NzR0FRVUZCd0VCQklJQlVUQ0NBVTB3ZWdZSUt3WUJCUVVITUFLRwpibWgwZEhBNkx5OTNZbU5qWVM1d2Eya3lMbk55ZGk1M1pYTjBjR0ZqTG1OdmJTNWhkUzlCU1VFdlYyVnpkSEJoWXlVeU1GTklRVElsCk1qQlRVMHdsTWpCRFFTVXlNRmRUUkVNdlYyVnpkSEJoWXlVeU1GTklRVElsTWpCVFUwd2xNakJEUVNVeU1GZFRSRU11WTNKME1JSE8KQmdnckJnRUZCUWN3QW9hQndXeGtZWEE2THk4dlEwNDlWMlZ6ZEhCaFl5VXlNRk5JUVRJbE1qQlRVMHdsTWpCRFFTVXlNRmRUUkVNcwpRMDQ5UVVsQkxFTk9QVkIxWW14cFl5VXlNRXRsZVNVeU1GTmxjblpwWTJWekxFTk9QVk5sY25acFkyVnpMRU5PUFVOdmJtWnBaM1Z5CllYUnBiMjRzUkVNOWQySmpZWFVzUkVNOWQyVnpkSEJoWXl4RVF6MWpiMjBzUkVNOVlYVS9ZMEZEWlhKMGFXWnBZMkYwWlQ5aVlYTmwKUDI5aWFtVmpkRU5zWVhOelBXTmxjblJwWm1sallYUnBiMjVCZFhSb2IzSnBkSGt3Q3dZRFZSMFBCQVFEQWdXZ01Ec0dDU3NHQVFRQgpnamNWQndRdU1Dd0dKQ3NHQVFRQmdqY1ZDSVhlc3lYZ2hRdUMwWTBwaDhyaGVzZjhFNEZZMStKS2hMcnBZUUlCWkFJQkNUQVRCZ05WCkhTVUVEREFLQmdnckJnRUZCUWNEQXpBYkJna3JCZ0VFQVlJM0ZRb0VEakFNTUFvR0NDc0dBUVVGQndNRE1JSUJLZ1lEVlIwZ0JJSUIKSVRDQ0FSMHdYUVlMS3dZQkJBR2NFNGRvQWdNd1RqQk1CZ2dyQmdFRkJRY0NBUlpBYUhSMGNEb3ZMM2RpWTJOaExuQnJhVEl1YzNKMgpMbmRsYzNSd1lXTXVZMjl0TG1GMUwxZGxjM1J3WVdOUWIyeHBZM2t2VjBKRFgwTlFVekl1Y0dSbUFEQmRCZ3dyQmdFRUFad1RoMmdCCkFRUXdUVEJMQmdnckJnRUZCUWNDQVJZL2FIUjBjRG92TDNkaVkyTmhMbkJyYVRJdWMzSjJMbmRsYzNSd1lXTXVZMjl0TG1GMUwxZGwKYzNSd1lXTlFiMnhwWTNrdlYwSkRYME5RTWk1d1pHWUFNRjBHRENzR0FRUUJuQk9IYUFFQkFUQk5NRXNHQ0NzR0FRVUZCd0lCRmo5bwpkSFJ3T2k4dmQySmpZMkV1Y0d0cExuTnlkaTUzWlhOMGNHRmpMbU52YlM1aGRTOVhaWE4wY0dGalVHOXNhV041TDFkQ1ExOUpWRk5RCkxuQmtaZ0F3RFFZSktvWklodmNOQVFFTEJRQURnZ0VCQUVPMnkwRFR3cGNuSDgvcURkVVpQK1JweGxvMm5sWVNhQmR3bnRBZWtuNkkKYXV5ZURhUEIxN2I3UkJVVzNNSU1zcWdMa2xSQWtKQTVsdTJMUy9YYldteXZvR2lGQnV4RTIxNktrbFNvWmlWeGRDR2p2M2ZJVGNpdgppVW9Mb1dQR2krSHZmelBEVXBIdkg5NUFBL0swOXBGM1ZCODJNRWc1dVNtTjkzMTQwdGp4eXFpQnJ5WEUvT2FRTmk3NEdoNWxWbU9iCkUzRnBnTEdSSGw2ZlFNeGx3UFRhYXpuZ0c5dTY2SU95ZWwvcVVwTkxzajlJY1U1U2VPTjJEMHg3bExVSWttcVM0cGloM2tSL3Q0OUUKQWdieDFEamw3akNIRVh6dzdxS3BCS1hGS2t1WkRLMDZ3R0laQmxsOGlLc2xQQ0E5MkZySllQbW05a1dTNzlXVm54WFNVUTQ9CjwvZHM6WDUwOUNlcnRpZmljYXRlPjwvZHM6WDUwOURhdGE+PC9kczpLZXlJbmZvPjwvZHM6U2lnbmF0dXJlPjxzYW1sOlN1YmplY3Q+PHNhbWw6TmFtZUlEIEZvcm1hdD0idXJuOm9hc2lzOm5hbWVzOnRjOlNBTUw6MS4xOm5hbWVpZC1mb3JtYXQ6dW5zcGVjaWZpZWQiIE5hbWVRdWFsaWZpZXI9Imh0dHBzOi8vbG9naW4uYXV0aGJyaWRnZS53ZXN0cGFjZ3JvdXAuY29tL25pZHAvc2FtbDIvbWV0YWRhdGEiIFNQTmFtZVF1YWxpZmllcj0idXJuOmFtYXpvbjp3ZWJzZXJ2aWNlcyI+TDEyMjg5Mzwvc2FtbDpOYW1lSUQ+PHNhbWw6U3ViamVjdENvbmZpcm1hdGlvbiBNZXRob2Q9InVybjpvYXNpczpuYW1lczp0YzpTQU1MOjIuMDpjbTpiZWFyZXIiPjxzYW1sOlN1YmplY3RDb25maXJtYXRpb25EYXRhIE5vdE9uT3JBZnRlcj0iMjAyMC0wMy0wNVQwMjoxMDoxOFoiIFJlY2lwaWVudD0iaHR0cHM6Ly9zaWduaW4uYXdzLmFtYXpvbi5jb20vc2FtbCIvPjwvc2FtbDpTdWJqZWN0Q29uZmlybWF0aW9uPjwvc2FtbDpTdWJqZWN0PjxzYW1sOkNvbmRpdGlvbnMgTm90QmVmb3JlPSIyMDIwLTAzLTA1VDAyOjAwOjE4WiIgTm90T25PckFmdGVyPSIyMDIwLTAzLTA1VDAyOjEwOjE4WiI+PHNhbWw6QXVkaWVuY2VSZXN0cmljdGlvbj48c2FtbDpBdWRpZW5jZT51cm46YW1hem9uOndlYnNlcnZpY2VzPC9zYW1sOkF1ZGllbmNlPjwvc2FtbDpBdWRpZW5jZVJlc3RyaWN0aW9uPjwvc2FtbDpDb25kaXRpb25zPjxzYW1sOkF1dGhuU3RhdGVtZW50IEF1dGhuSW5zdGFudD0iMjAyMC0wMy0wNVQwMTo0MToyNFoiIFNlc3Npb25JbmRleD0iaWRWTzRuRkVNbm9TNFJfemRaNVZNX1RlcGdmVGciPjxzYW1sOkF1dGhuQ29udGV4dD48c2FtbDpBdXRobkNvbnRleHRDbGFzc1JlZj51cm46b2FzaXM6bmFtZXM6dGM6U0FNTDoyLjA6YWM6Y2xhc3NlczpLZXJiZXJvczwvc2FtbDpBdXRobkNvbnRleHRDbGFzc1JlZj48c2FtbDpBdXRobkNvbnRleHREZWNsUmVmPmF3cy9tZmEvYXV0aC91cmk8L3NhbWw6QXV0aG5Db250ZXh0RGVjbFJlZj48L3NhbWw6QXV0aG5Db250ZXh0Pjwvc2FtbDpBdXRoblN0YXRlbWVudD48c2FtbDpBdHRyaWJ1dGVTdGF0ZW1lbnQ+PHNhbWw6QXR0cmlidXRlIHhtbG5zOnhzPSJodHRwOi8vd3d3LnczLm9yZy8yMDAxL1hNTFNjaGVtYSIgeG1sbnM6eHNpPSJodHRwOi8vd3d3LnczLm9yZy8yMDAxL1hNTFNjaGVtYS1pbnN0YW5jZSIgTmFtZT0iaHR0cHM6Ly9hd3MuYW1hem9uLmNvbS9TQU1ML0F0dHJpYnV0ZXMvUm9sZVNlc3Npb25OYW1lIiBOYW1lRm9ybWF0PSJ1cm46b2FzaXM6bmFtZXM6dGM6U0FNTDoyLjA6YXR0cm5hbWUtZm9ybWF0OnVuc3BlY2lmaWVkIj48c2FtbDpBdHRyaWJ1dGVWYWx1ZSB4c2k6dHlwZT0ieHM6c3RyaW5nIj5MMTIyODkzPC9zYW1sOkF0dHJpYnV0ZVZhbHVlPjwvc2FtbDpBdHRyaWJ1dGU+PHNhbWw6QXR0cmlidXRlIHhtbG5zOnhzPSJodHRwOi8vd3d3LnczLm9yZy8yMDAxL1hNTFNjaGVtYSIgeG1sbnM6eHNpPSJodHRwOi8vd3d3LnczLm9yZy8yMDAxL1hNTFNjaGVtYS1pbnN0YW5jZSIgTmFtZT0iaHR0cHM6Ly9hd3MuYW1hem9uLmNvbS9TQU1ML0F0dHJpYnV0ZXMvUm9sZSIgTmFtZUZvcm1hdD0idXJuOm9hc2lzOm5hbWVzOnRjOlNBTUw6Mi4wOmF0dHJuYW1lLWZvcm1hdDp1bnNwZWNpZmllZCI+PHNhbWw6QXR0cmlidXRlVmFsdWUgeHNpOnR5cGU9InhzOnN0cmluZyI+YXJuOmF3czppYW06Ojg2MzM2OTQzMDA1NTpyb2xlL3diYy1hZG1pbixhcm46YXdzOmlhbTo6ODYzMzY5NDMwMDU1OnNhbWwtcHJvdmlkZXIvd2VzdHBhY2lkcDwvc2FtbDpBdHRyaWJ1dGVWYWx1ZT48c2FtbDpBdHRyaWJ1dGVWYWx1ZSB4c2k6dHlwZT0ieHM6c3RyaW5nIj5hcm46YXdzOmlhbTo6ODYzMzY5NDMwMDU1OnJvbGUvd2JjLWVuZ2luZWVyLGFybjphd3M6aWFtOjo4NjMzNjk0MzAwNTU6c2FtbC1wcm92aWRlci93ZXN0cGFjaWRwPC9zYW1sOkF0dHJpYnV0ZVZhbHVlPjxzYW1sOkF0dHJpYnV0ZVZhbHVlIHhzaTp0eXBlPSJ4czpzdHJpbmciPmFybjphd3M6aWFtOjo4NjMzNjk0MzAwNTU6cm9sZS93YmMtbWFuYWdlcixhcm46YXdzOmlhbTo6ODYzMzY5NDMwMDU1OnNhbWwtcHJvdmlkZXIvd2VzdHBhY2lkcDwvc2FtbDpBdHRyaWJ1dGVWYWx1ZT48c2FtbDpBdHRyaWJ1dGVWYWx1ZSB4c2k6dHlwZT0ieHM6c3RyaW5nIj5hcm46YXdzOmlhbTo6ODYzMzY5NDMwMDU1OnJvbGUvd2JjLXZpZXdvbmx5LGFybjphd3M6aWFtOjo4NjMzNjk0MzAwNTU6c2FtbC1wcm92aWRlci93ZXN0cGFjaWRwPC9zYW1sOkF0dHJpYnV0ZVZhbHVlPjxzYW1sOkF0dHJpYnV0ZVZhbHVlIHhzaTp0eXBlPSJ4czpzdHJpbmciPmFybjphd3M6aWFtOjo4NjMzNjk0MzAwNTU6cm9sZS93YmMtcmVhZG9ubHksYXJuOmF3czppYW06Ojg2MzM2OTQzMDA1NTpzYW1sLXByb3ZpZGVyL3dlc3RwYWNpZHA8L3NhbWw6QXR0cmlidXRlVmFsdWU+PHNhbWw6QXR0cmlidXRlVmFsdWUgeHNpOnR5cGU9InhzOnN0cmluZyI+YXJuOmF3czppYW06OjE2OTI2NzQ1MTI4NTpyb2xlL3diYy1hZG1pbixhcm46YXdzOmlhbTo6MTY5MjY3NDUxMjg1OnNhbWwtcHJvdmlkZXIvd2VzdHBhY2lkcDwvc2FtbDpBdHRyaWJ1dGVWYWx1ZT48c2FtbDpBdHRyaWJ1dGVWYWx1ZSB4c2k6dHlwZT0ieHM6c3RyaW5nIj5hcm46YXdzOmlhbTo6MTY5MjY3NDUxMjg1OnJvbGUvd2JjLWVuZ2luZWVyLGFybjphd3M6aWFtOjoxNjkyNjc0NTEyODU6c2FtbC1wcm92aWRlci93ZXN0cGFjaWRwPC9zYW1sOkF0dHJpYnV0ZVZhbHVlPjxzYW1sOkF0dHJpYnV0ZVZhbHVlIHhzaTp0eXBlPSJ4czpzdHJpbmciPmFybjphd3M6aWFtOjoxNjkyNjc0NTEyODU6cm9sZS93YmMtdmlld29ubHksYXJuOmF3czppYW06OjE2OTI2NzQ1MTI4NTpzYW1sLXByb3ZpZGVyL3dlc3RwYWNpZHA8L3NhbWw6QXR0cmlidXRlVmFsdWU+PHNhbWw6QXR0cmlidXRlVmFsdWUgeHNpOnR5cGU9InhzOnN0cmluZyI+YXJuOmF3czppYW06OjQwMTI1MjgzMDY4Njpyb2xlL3diYy1hZG1pbixhcm46YXdzOmlhbTo6NDAxMjUyODMwNjg2OnNhbWwtcHJvdmlkZXIvd2VzdHBhY2lkcDwvc2FtbDpBdHRyaWJ1dGVWYWx1ZT48c2FtbDpBdHRyaWJ1dGVWYWx1ZSB4c2k6dHlwZT0ieHM6c3RyaW5nIj5hcm46YXdzOmlhbTo6NDAxMjUyODMwNjg2OnJvbGUvd2JjLWVuZ2luZWVyLGFybjphd3M6aWFtOjo0MDEyNTI4MzA2ODY6c2FtbC1wcm92aWRlci93ZXN0cGFjaWRwPC9zYW1sOkF0dHJpYnV0ZVZhbHVlPjxzYW1sOkF0dHJpYnV0ZVZhbHVlIHhzaTp0eXBlPSJ4czpzdHJpbmciPmFybjphd3M6aWFtOjo0MDEyNTI4MzA2ODY6cm9sZS93YmMtdmlld29ubHksYXJuOmF3czppYW06OjQwMTI1MjgzMDY4NjpzYW1sLXByb3ZpZGVyL3dlc3RwYWNpZHA8L3NhbWw6QXR0cmlidXRlVmFsdWU+PHNhbWw6QXR0cmlidXRlVmFsdWUgeHNpOnR5cGU9InhzOnN0cmluZyI+YXJuOmF3czppYW06OjQwMTI1MjgzMDY4Njpyb2xlL3diYy1yZWFkb25seSxhcm46YXdzOmlhbTo6NDAxMjUyODMwNjg2OnNhbWwtcHJvdmlkZXIvd2VzdHBhY2lkcDwvc2FtbDpBdHRyaWJ1dGVWYWx1ZT48c2FtbDpBdHRyaWJ1dGVWYWx1ZSB4c2k6dHlwZT0ieHM6c3RyaW5nIj5hcm46YXdzOmlhbTo6MTY5MjY3NDUxMjg1OnJvbGUvd2JjLXJlYWRvbmx5LGFybjphd3M6aWFtOjoxNjkyNjc0NTEyODU6c2FtbC1wcm92aWRlci93ZXN0cGFjaWRwPC9zYW1sOkF0dHJpYnV0ZVZhbHVlPjwvc2FtbDpBdHRyaWJ1dGU+PC9zYW1sOkF0dHJpYnV0ZVN0YXRlbWVudD48L3NhbWw6QXNzZXJ0aW9uPjwvc2FtbHA6UmVzcG9uc2U+"

	//when
	samlResponseDoc, err := goquery.NewDocumentFromReader(bytes.NewReader(samlResponseData))
	require.Nil(t, err)
	actualResult, err := extractSAMLAssertion(samlResponseDoc)
	require.Nil(t, err)

	//then
	require.Equal(t, expectedResult, actualResult)
}

func TestExtractGetToContentUrlPositive(t *testing.T) {
	//given
	getToContentData, err := ioutil.ReadFile("responses/getToContent.html")
	require.Nil(t, err)
	expectedResourceUrl := "/nidp/jsp/content.jsp?sid=0&option=credential&id=AWS"

	//when
	getToContentDoc, err := goquery.NewDocumentFromReader(bytes.NewReader(getToContentData))
	require.Nil(t, err)
	actualResourceUrl, ok := extractGetToContentUrl(getToContentDoc)

	//then
	require.True(t, ok)
	require.Equal(t, expectedResourceUrl, actualResourceUrl)
}

func TestExtractGetToContentUrlNegative(t *testing.T) {
	//given
	samlResposeData, err := ioutil.ReadFile("responses/samlRespose.html")
	require.Nil(t, err)

	//when
	samlResposeDoc, err := goquery.NewDocumentFromReader(bytes.NewReader(samlResposeData))
	require.Nil(t, err)
	actualResourceUrl, ok := extractGetToContentUrl(samlResposeDoc)

	//then
	require.False(t, ok)
	require.Equal(t, "", actualResourceUrl)
}

func TestExtractGetToContentUrlDiv(t *testing.T) {
	//given
	getToContentData, err := ioutil.ReadFile("responses/getToContentDiv.html")
	require.Nil(t, err)
	expectedResourceUrl := "/nidp/app/login?id=contract_kerb&sid=0&option=credential&sid=0"

	//when
	getToContentDoc, err := goquery.NewDocumentFromReader(bytes.NewReader(getToContentData))
	require.Nil(t, err)
	actualResourceUrl, ok := extractGetToContentUrl(getToContentDoc)

	//then
	require.True(t, ok)
	require.Equal(t, expectedResourceUrl, actualResourceUrl)
}

func TestExtractWinLocHrefUrlPositive(t *testing.T) {
	//given
	winLocHrefData, err := ioutil.ReadFile("responses/winLocHref.html")
	require.Nil(t, err)
	expectedResourceUrl := "https://login.authbridge.somegroup.com/nidp/saml2/idpsend?PID=STSPv8a5kc"

	//when
	winLocHrefDoc, err := goquery.NewDocumentFromReader(bytes.NewReader(winLocHrefData))
	require.Nil(t, err)
	actualResourceUrl, ok := extractWinLocHrefURL(winLocHrefDoc)

	//then
	require.True(t, ok)
	require.Equal(t, expectedResourceUrl, actualResourceUrl)
}

func TestExtractWinLocHrefUrlNegative(t *testing.T) {
	//given
	samlResposeData, err := ioutil.ReadFile("responses/samlRespose.html")
	require.Nil(t, err)

	//when
	samlResposeDoc, err := goquery.NewDocumentFromReader(bytes.NewReader(samlResposeData))
	require.Nil(t, err)
	actualResourceUrl, ok := extractWinLocHrefURL(samlResposeDoc)

	//then
	require.False(t, ok)
	require.Equal(t, "", actualResourceUrl)
}

func TestExtractIDPLoginPassPositive(t *testing.T) {
	//given
	idpLoginPassData, err := ioutil.ReadFile("responses/idpLoginPass.html")
	require.Nil(t, err)
	expectedForm := &page.Form{
		URL:    "https://login.authbridge.somegroup.com/nidp/app/login?sid=0&sid=0",
		Method: "POST",
		Values: &url.Values{},
	}

	//when
	idpLoginPassDoc, err := goquery.NewDocumentFromReader(bytes.NewReader(idpLoginPassData))
	require.Nil(t, err)
	actualForm, ok := extractIDPLoginPass(idpLoginPassDoc)

	//then
	require.True(t, ok)
	require.Equal(t, expectedForm, actualForm)
}

func TestExtractIDPLoginPassNegative(t *testing.T) {
	//given
	idpLoginRsaData, err := ioutil.ReadFile("responses/idpLoginRsa.html")
	require.Nil(t, err)

	//when
	idpLoginRsaDoc, err := goquery.NewDocumentFromReader(bytes.NewReader(idpLoginRsaData))
	require.Nil(t, err)
	actualResourceUrl, ok := extractIDPLoginPass(idpLoginRsaDoc)

	//then
	require.False(t, ok)
	require.Nil(t, actualResourceUrl)
}

func TestExtractPrivilegedIDPLoginPassPositive(t *testing.T) {
	//given
	idpLoginPassData, err := ioutil.ReadFile("responses/privileged_flow/idpLoginPass.html")
	require.Nil(t, err)
	expectedForm := &page.Form{
		URL:    "https://login.authbridge.somegroup.com/nidp/app/login?sid=0&sid=0",
		Method: "POST",
		Values: &url.Values{},
	}

	//when
	idpLoginPassDoc, err := goquery.NewDocumentFromReader(bytes.NewReader(idpLoginPassData))
	require.Nil(t, err)
	actualForm, ok := extractIDPLoginPass(idpLoginPassDoc)

	//then
	require.True(t, ok)
	require.Equal(t, expectedForm, actualForm)
}

func TestExtractIDPLoginRsaPositive(t *testing.T) {
	//given
	idpLoginRsaData, err := ioutil.ReadFile("responses/idpLoginRsa.html")
	require.Nil(t, err)
	expectedForm := &page.Form{
		URL:    "https://login.authbridge.somegroup.com/nidp/app/login?sid=11&sid=11",
		Method: "POST",
		Values: &url.Values{},
	}

	//when
	idpLoginRsaDoc, err := goquery.NewDocumentFromReader(bytes.NewReader(idpLoginRsaData))
	require.Nil(t, err)
	actualForm, ok := extractIDPLoginRsa(idpLoginRsaDoc)

	//then
	require.True(t, ok)
	require.Equal(t, expectedForm, actualForm)
}

func TestExtractIDPLoginRsaNegative(t *testing.T) {
	//given
	idpLoginPassData, err := ioutil.ReadFile("responses/idpLoginPass.html")
	require.Nil(t, err)

	//when
	idpLoginPassDoc, err := goquery.NewDocumentFromReader(bytes.NewReader(idpLoginPassData))
	require.Nil(t, err)
	actualResourceUrl, ok := extractIDPLoginRsa(idpLoginPassDoc)

	//then
	require.False(t, ok)
	require.Nil(t, actualResourceUrl)
}

func TestExtractPrivilegedIDPLoginRsaNegative(t *testing.T) {
	//given
	idpLoginPassData, err := ioutil.ReadFile("responses/privileged_flow/idpLoginPass.html")
	require.Nil(t, err)

	//when
	idpLoginPassDoc, err := goquery.NewDocumentFromReader(bytes.NewReader(idpLoginPassData))
	require.Nil(t, err)
	actualResourceUrl, ok := extractIDPLoginRsa(idpLoginPassDoc)

	//then
	require.False(t, ok)
	require.Nil(t, actualResourceUrl)
}

func TestPrivilegedLoginUrl(t *testing.T) {
	//given
	mfa := "Privileged"
	baseUrl := "https://abc.com"
	defaultResourcePath := "/login.html"
	expectedLoginUrl := baseUrl + "/nidp/app/login?id=privacc&sid=0&option=credential"

	//when
	loginUrl, err := getLoginUrl(mfa, baseUrl, defaultResourcePath)

	//then
	require.Nil(t, err)
	require.Equal(t, loginUrl, expectedLoginUrl)
}

func TestDefaultLoginUrl(t *testing.T) {
	//given
	mfa := "Auto"
	baseUrl := "https://abc.com"
	defaultResourcePath := "/login.html"
	expectedLoginUrl := baseUrl + defaultResourcePath

	//when
	loginUrl, err := getLoginUrl(mfa, baseUrl, defaultResourcePath)

	//then
	require.Nil(t, err)
	require.Equal(t, loginUrl, expectedLoginUrl)
}

func TestUnsupportedMFA(t *testing.T) {
	//given
	mfa := "None"
	baseUrl := "https://abc.com"
	defaultResourcePath := "/login.html"
	expectedErrorString := "Unsupported MFA"

	//when
	_, err := getLoginUrl(mfa, baseUrl, defaultResourcePath)

	//then
	require.EqualError(t, err, expectedErrorString)
}
