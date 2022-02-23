package si

import (
	"fmt"
	"sync"
)

/*
疑問:
ネット上の例を見ると、シングルトンの対象となる構造体がパブリックで、かつ同じパッケージに書かれている事が多い。
その場合、普通にコンストラクタを使わずに構造体を初期化できてしまうのでは？
それはシングルトンを実現できていると言えるのか？
以下のように構造体をプライベートにして別パッケージに切り分けないといけないのではないか？
*/

type singleton struct{}

var singletonInstance singleton
var once sync.Once

func GetSingletonInstance() singleton {
	once.Do(func() {
		fmt.Println("インスタンスを生成しました。")
		singletonInstance = singleton{}
	})
	return singletonInstance
}
