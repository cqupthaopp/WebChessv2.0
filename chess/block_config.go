/**
 * @Author: Hao_pp
 * @Data: 2022年7月17日-16点03分
 * @Desc: 无无无
 */

package chess

var BlockFlag map[string]int
var BlockNumFlag map[int]string

func InitBlock() {

	{

		BlockFlag = map[string]int{
			"帅": 1,
			"士": 2,
			"象": 3,
			"马": 4,
			"车": 5,
			"炮": 6,
			"兵": 7,
		}

		BlockNumFlag = make(map[int]string, 7)
		for key, value := range BlockFlag {
			BlockNumFlag[value] = key
		}

	} //棋子的映射

}
