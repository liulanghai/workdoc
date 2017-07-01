//下载中心接口
type url_down_center int

type Store_file_in struct {
	file_md5        string
	file_store_path string //文件在下载中心的存储位置，全路径
}
type Store_file_out struct {
	status int
}

type Get_url_in struct {
	url string
}
type Get_url_out struct {
	status          int
	file_md5        string
	file_name       string
	file_size       string
	file_store_path string
}

func (r *url_down_center) Store_file(in Store_file_in, reply *Store_file_out) error {
	
}
func (r *url_down_center) Get_url(in Get_url_in, reply *Get_url_out) error {
	
}
/*************************************************/

//url查杀服务接口
type url_kill_server int

type Url_kill_in struct {
	url string
}
type Url_kill_out struct {
	url           string
	status        int
	danger_detect string
}
func (r *url_kill_server) Url_kill(in Url_kill_in, reply *Url_kill_out) error {
	
}
/*************************************************/
//url鉴定中心接口
type url_analysis_server int

type Url_analysis_in struct {
	url string
}
type Url_analysis_out struct {
	url           string
	status        int
	danger_detect string
	report_id     string
}
func (r *url_analysis_server) Url_analysis(in Url_analysis_in, reply *Url_analysis_out) error {
	
}
//文件鉴定中心接口
type file_analysis_server int

type File_analysis_in struct {
	file_md5 string
	file_url string
}

type File_analysis_out struct {
	file_md5      string
	file_sha256   string
	status        int
	danger_detect string
	report_id     string
}

//鉴定日志中心接口
type Analysis_log_centure int 
type Analysis_
























