package setting

import "github.com/spf13/viper"

type Setting struct {
	vp *viper.Viper
}

func NewSetting() (*Setting, error) {
	vp := viper.New()
	vp.SetConfigName("config")
	vp.AddConfigPath("configs/")
	vp.SetConfigType("yaml")
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}
	// 另一個寫法就直接讀進變數裡
	// var config Config
	// if err := viper.Unmarshal(&config); err != nil {
	// 		fmt.Println(err)
	// 		return nil, err
	// }
	return &Setting{vp}, nil
}
