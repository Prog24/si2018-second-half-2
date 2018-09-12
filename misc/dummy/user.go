package main

import (
	"math/rand"
	"time"

	"github.com/go-openapi/strfmt"

	"github.com/eure/si2018-server-side/entities"
	"github.com/eure/si2018-server-side/repositories"
)

func dummyUser() {
	state := []string{
		"北海道", "青森", "岩手", "宮城", "秋田", "山形", "福島", "茨城", "栃木", "群馬", "埼玉",
		"千葉", "東京", "神奈川", "新潟", "富山", "石川", "福井", "山梨", "長野", "岐阜", "静岡",
		"愛知", "三重", "滋賀", "京都", "大阪", "兵庫", "奈良", "和歌山", "鳥取", "島根", "岡山",
		"広島", "山口", "徳島", "香川", "愛媛", "高知", "福岡", "佐賀", "長崎", "熊本", "大分",
		"宮崎", "鹿児島", "沖縄", "海外", "宇宙",
	}

	maleNickname := []string{"アキト", "アキヒコ", "アキヒロ", "アキラ", "アサヒ", "アツキ", "アツシ", "アツヤ", "アツロウ", "アマネ", "アユム", "イオリ",
		"イズモ", "イチロウ", "イツキ", "イブキ", "イマハル", "ウキョウ", "ウコン", "エイジ", "エツシ", "オウスケ", "オオヒト", "オトヒコ", "カイ", "カイト",
		"カケル", "カズキ", "カズトシ", "カズマ", "カズヤ", "カズユキ", "カンキチ", "カンスケ", "カンタ", "カンチ", "ギイチロウ", "キョウスケ", "キョウヘイ",
		"キヨタケ", "キンペイ", "ケイゴ", "ケイジ", "ケイスケ", "ケイタツ", "ケン", "ゲンカイ", "ゲンキ", "ケンゴ", "ケンスケ", "ケンタ", "ケンタロウ", "ケント",
		"ケンヤ", "コウ", "ゴウ", "コウイチ", "コウイチロウ", "コウキ", "コウジ", "コウスケ", "コウセイ", "コウタ", "コウタロウ", "コウヘイ", "コウヤ", "コウヨウ",
		"サイト", "サトシ", "サトル", "サモン", "シュウ", "シュウト", "シュウヘイ", "シュウヤ", "シュン", "ジュン", "シュンスケ", "シュンペイ", "ジュンペイ", "シュンヤ",
		"ジュンヤ", "ショウ", "ショウゴ", "ショウタ", "ショウタロウ", "ショウヘイ", "ショウマ", "ショウヤ", "ジン", "シンイチ", "シンゴ", "シンジ", "シンタロウ",
		"シンノスケ", "シンペイ", "シンヤ", "スオウ", "スグル", "セイヤ", "センイチ", "ソウイチロウ", "ソウスケ", "ソウマ", "ソラ", "タイキ", "ダイキ", "ダイゴ",
		"ダイスケ", "タイセイ", "タイソン", "ダイチ", "タイチ", "タカシ", "タカヒロ", "タカヤ", "タカユキ", "タク", "タクマ", "タクミ", "タクヤ", "タクロウ",
		"タケシ", "タケル", "タスク", "タツキ", "タツヤ", "タロウ", "ツカサ", "ツバサ", "ツヨシ", "ツルギ", "テッペイ", "テツヤ", "テンドウ", "トウゴ", "トウマ",
		"トウヨウ", "トカチ", "トオル", "トシキ", "トシヤ", "トミヨシ", "トモキ", "トモヒロ", "トモヤ", "トヨアキ", "ナオキ", "ナオト", "ナオヤ", "ナガト", "ナツヒコ",
		"ナンヨウ", "ノブアキ", "ノブオ", "ハジメ", "ハヤト", "ハルキ", "ハルト", "ヒデオ", "ヒビキ", "ヒュウガ", "ヒリュウ", "ヒロ", "ヒロアキ", "ヒロキ", "ヒロシ",
		"ヒロタカ", "ヒロト", "ヒロム", "ヒロユキ", "フウタ", "フミヤ", "ホクト", "ホダカ", "マコト", "マサキ", "マサシ", "マサタカ", "マサト", "マサノブ", "マサヒロ",
		"マサヤ", "マサユキ", "マサル", "マリオ", "ミツル", "ミノル", "ムサシ", "モトキ", "ヤスヒロ", "ヤマト", "ユウ", "ユウイチ", "ユウイチロウ", "ユウゴ", "ユウサク",
		"ユウジ", "ユウスケ", "ユウダイ", "ユウタロウ", "ユウト", "ユウヘイ", "ユウマ", "ユウヤ", "ユウリ", "ユタカ", "ヨウスケ", "ヨウヘイ", "ヨシモリ", "リク",
		"リクト", "リュウ", "リュウキ", "リュウジ", "リュウスケ", "リュウタ", "リュウタロウ", "リュウト", "リュウスケ", "リュウノスケ", "リュウヤ", "リョウ", "リョウガ",
		"リョウジ", "リョウタ", "リョウヘイ", "リョウマ", "リンタロウ", "ヨウスケ", "ルイ", "レブン", "レン", "ワタル"}

	femaleNickname := []string{"アイ", "アイカナ", "アイリ", "アオイ", "アカネ", "アカリ", "アキコ", "アキホ", "アクリ", "アケミ", "アスカ", "アミ", "アヤ",
		"アヤカ", "アヤナ", "アヤネ", "アヤノ", "アヤメ", "アユ", "アユミ", "アリサ", "アリス", "アンナ", "イズミ", "ウタ", "ウヅキ", "ウメコ", "エナ", "エミ",
		"エミリ", "エリ", "エリカ", "エリナ", "カエデ", "カオリ", "カオル", "カザネ", "カナ", "カナエ", "カナコ", "カナミ", "カノコ", "カノン", "カリン", "カレン",
		"カンナ", "キミ", "キヨ", "キョウカ", "キョウコ", "キヨコ", "クルミ", "ケイ", "ケイコ", "コトネ", "コトノ", "コトミ", "コノハ", "コハル", "コマチ", "コユキ",
		"サアヤ", "サエ", "サオリ", "サキ", "サチコ", "サツキ", "サトミ", "サナ", "サヤ", "サヤカ", "サユリ", "サラ", "シオリ", "シオン", "シノ", "シノブ", "シホ",
		"ジュリア", "スズ", "スズカ", "スズネ", "スミレ", "セイ", "セイカ", "セイラ", "セナ", "セリカ", "セリナ", "センリ", "ソノコ", "タカコ", "タマキ", "チアキ",
		"チカ", "チサト", "チナツ", "チナミ", "チヨ", "テルコ", "トシコ", "トモカ", "トモコ", "トモエ", "トモハ", "トモミ", "ナオ", "ナオコ", "ナギサ", "ナツキ",
		"ナナコ", "ナナセ", "ナナミ", "ナホ", "ナミ", "ナルミ", "ノゾミ", "ノドカ", "ハヅキ", "ハツミ", "ハナ", "ハルカ", "ヒカリ", "ヒカル", "ヒトミ", "ヒナ",
		"ヒナコ", "ヒナタ", "ヒナノ", "ヒロミ", "フウカ", "フウコ", "フミカ", "フミコ", "ホタル", "ホナミ", "ホノカ", "マイカ", "マイコ", "マオ", "マキ", "マサコ",
		"マドカ", "マナカ", "マナミ", "マヒロ", "マホ", "マミ", "マヤ", "マユコ", "マリ", "マリカ", "マリナ", "マリン", "ミウ", "ミオ", "ミカ", "ミキ", "ミサ",
		"ミサト", "ミズキ", "ミスズ", "ミズホ", "ミツ", "ミツキ", "ミヅキ", "ミナ", "ミナミ", "ミノリ", "ミホ", "ミユ", "ミユウ", "ミユキ", "ミライ", "メイ",
		"メグミ", "モエカ", "モエコ", "モモ", "モモカ", "ヤワラ", "ユイ", "ユイカ", "ユウ", "ユウカ", "ユウコ", "ユウナ", "ユカリ", "ユキ", "ユキナ", "ユキノ",
		"ユナ", "ユミ", "ユメ", "ユメカ", "ユリ", "ユリカ", "ユリエ", "ユリナ", "ヨシコ", "ヨシノ", "リエ", "リオ", "リカ", "リカコ", "リコ", "リサ", "リサコ",
		"リナ", "リノ", "リホ", "リョウコ", "リリカ", "リンカ", "ルカ", "ルナ", "レイ", "レイカ", "レイナ", "レオ", "レナ", "ワカナ"}

	child := []string{
		"いない",
		"同居中",
		"別居中",
	}

	job := []string{
		"会社員", "医師", "弁護士", "公認会計士", "経営者・役員", "公務員", "事務員", "大手商社", "外資金融", "大手企業",
		"大手外資", "マスコミ・広告", "クリエイター", "IT関連", "パイロット", "客室乗務員", "芸能・モデル",
		"アパレル・ショップ", "アナウンサー", "イベントコンパニオン", "受付", "秘書", "看護師", "保育士",
		"自由業", "学生", "その他", "上場企業", "金融", "コンサル", "調理師・栄養士", "教育関連", "食品関連",
		"製薬", "保険", "不動産", "建築関連", "通信", "流通", "WEB業界", "接客業", "美容関係", "エンターテインメント",
		"旅行関係", "ブライダル", "福祉・介護", "広告", "マスコミ", "薬剤師", "スポーツ選手",
	}

	education := []string{
		"短大/専門学校卒",
		"高校卒",
		"大学卒",
		"大学院卒",
	}

	income := []string{
		"200万円未満",
		"200万円以上〜400万円未満",
		"400万円以上〜600万円未満",
		"600万円以上〜800万円未満",
		"800万円以上〜1000万円未満",
		"1000万円以上〜1500万円未満",
		"1500万円以上〜2000万円未満",
		"2000万円以上〜3000万円未満",
		"3000万円以上",
	}

	tweet := []string{
		"肉食べたい",
		"寿司食べたいです",
		"スポーツしたいなあ",
		"お酒飲みたい！",
		"クラブにいきませんか",
		"コーヒーが好きです",
		"遊びに行きたいなあ",
	}

	introduction := []string{
		"はじめまして！●●在住の××歳です。プロフィール見てくれて、ありがとうございます！普段は、△△の仕事をしています。たまたま、facebookを見ていたらこのPairsを見つけて登録してみました！休日はみんなで◎◎をしたり、新しい場所に出掛けるのが好きで、天気のいい日はふらっと買い物に出かけたりしてます。恋愛に関しては、□□なタイプで、◆◆には弱いです。笑普段出会えない方とめぐり逢えたらと思っています。もし、少しでも気になって頂けたらよろしくお願いします！！",
		"はじめまして！●●在住の××歳です。プロフィールを見ていただきありがとうございます。いい出会いがあればいいな、と思って登録しました。性格は◇◇で、いつも友達に●●な奴だな〜と言われています。笑音楽が好きで、△△なんかを聞いて、気晴らしをしています！新しいことにチャレンジするのも好きで最近は■■をはじめてみています。もし、詳しい方がいらっしゃったら教えてください。一緒に楽しく過ごせる方と恋愛がしたいです！少しでも興味を持っていただければうれしいです！よろしくお願いします。",
		"こんばんは！facebookでこのサイトを見つけて登録してみました！●●在住の××歳です。△△の仕事をしています。趣味は、××と◇◇（例えば、映画）です。好きな◇◇（例えば、映画）は■■です。正直、仕事もそれなりに忙しく、新しい出会いもそう頻繁にはありません。せっかくなので、趣味の話が合う人と出会いたいです。少しでも共通の趣味がある方はぜひいいね！してください。よろしくお願いします。",
		"プロフィール読んでいただき、ありがとうございます！はじめまして！●●住まいの××歳です。出身は△△です。学生時代から◇◇が好きで、今も週末は◇◇を続けています。また、子どもと遊ぶのも好きで、友人の家で、子どもとよく戯れています。笑最近、周りでも結婚が増えてきて、少し意識し始めています。もちろん、相手ありきですが、大切な時間を一緒に過ごしていく人とめぐり逢いたいなと思っています。よろしくお願いします。",
		"こんにちは！はじめまして。●●の××歳、△△な仕事をしています。◇◇が好きで、比較的■■（アウトドアorインドア）なことが好きです！最近なかなかできないですが、時間があれば◇◇をしたいです。旅行も好きですが、最近は遠くには行けていません。◎◎（温泉など）でゆっくりするのが落ち着けていいですね。気が合う人とお話をしてみたいので、少しでも気になったら、いいね！してもらえるとうれしいです。よろしくお願いします。",
	}

	bodyBuild := []string{
		"スリム",
		"やや細め",
		"普通",
		"グラマー",
		"筋肉質",
		"ややぽっちゃり",
		"ぽっちゃり",
	}

	maritalStatus := []string{
		"独身(未婚)",
		"独身(離婚)",
		"独身(死別)",
	}

	whenMarry := []string{
		"すぐにでもしたい",
		"２〜３年のうちに",
		"良い人がいればしたい",
		"今のところ結婚は考えていない",
		"わからない",
	}

	wantChild := []string{
		"いいえ",
		"はい",
		"わからない",
		"相手と相談して決める",
	}

	smoking := []string{
		"吸わない",
		"吸う",
		"ときどき吸う",
		"非喫煙者の前では吸わない",
		"相手が嫌ならやめる",
		"吸う（電子タバコ）",
	}

	drinking := []string{
		"未設定",
		"飲まない",
		"飲む",
		"ときどき飲む",
	}

	holiday := []string{
		"土日",
		"平日",
		"不定期",
	}

	howToMeet := []string{
		"マッチング後にまずは会いたい",
		"気が合えば会いたい",
		"メッセージを重ねてから会いたい",
	}

	costOfDate := []string{
		"割り勘がいい",
		"おごってほしい",
		"おごりたい",
		"多めに出してほしい",
		"多めに出します",
		"相談して決めたい",
	}

	maleNthChild := []string{
		"長男",
		"次男",
		"三男",
		"一人っ子",
	}

	femaleNthChild := []string{
		"長女",
		"次女",
		"三女",
		"一人っ子",
	}

	houseWork := []string{
		"積極的に参加したい",
		"できれば参加したい",
		"できれば相手に任せたい",
		"相手に任せたい",
	}

	height := []string{
		"150cm",
		"160cm",
		"170cm",
	}

	r := repositories.NewUserRepository()

	for i := maleIDStart; i <= maleIDEnd; i++ {
		rand.Seed(time.Now().UnixNano())

		createdDaysAgo := rand.Intn(600)
		updatedDaysAgo := createdDaysAgo / (rand.Intn(600) + 1)
		minute1 := rand.Intn(1440)
		minute2 := rand.Intn(1440)
		birthday := strfmt.Date(time.Now().Add((-time.Duration(20+rand.Intn(40))*365 + time.Duration(rand.Intn(365))) * (24 * time.Hour)))

		u := entities.User{
			ID:             int64(i),
			Gender:         "M",
			Nickname:       maleNickname[rand.Intn(len(maleNickname))],
			NthChild:       maleNthChild[rand.Intn(len(maleNthChild))],
			Birthday:       birthday,
			Tweet:          tweet[rand.Intn(len(tweet))],
			Introduction:   introduction[rand.Intn(len(introduction))],
			ResidenceState: state[rand.Intn(len(state))],
			HomeState:      state[rand.Intn(len(state))],
			Education:      education[rand.Intn(len(education))],
			Job:            job[rand.Intn(len(job))],
			AnnualIncome:   income[rand.Intn(len(income))],
			Height:         height[rand.Intn(len(height))],
			BodyBuild:      bodyBuild[rand.Intn(len(bodyBuild))],
			MaritalStatus:  maritalStatus[rand.Intn(len(maritalStatus))],
			Child:          child[rand.Intn(len(child))],
			WhenMarry:      whenMarry[rand.Intn(len(whenMarry))],
			WantChild:      wantChild[rand.Intn(len(wantChild))],
			Smoking:        smoking[rand.Intn(len(smoking))],
			Drinking:       drinking[rand.Intn(len(drinking))],
			Holiday:        holiday[rand.Intn(len(holiday))],
			HowToMeet:      howToMeet[rand.Intn(len(howToMeet))],
			CostOfDate:     costOfDate[rand.Intn(len(costOfDate))],
			Housework:      houseWork[rand.Intn(len(houseWork))],
			CreatedAt:      strfmt.DateTime(time.Now().AddDate(0, 0, -createdDaysAgo).Add(-time.Duration(minute1) * time.Minute)),
			UpdatedAt:      strfmt.DateTime(time.Now().AddDate(0, 0, -updatedDaysAgo).Add(-time.Duration(minute2) * time.Minute)),
		}
		r.Create(u)
	}

	for i := femaleIDStart; i <= femaleIDEnd; i++ {
		rand.Seed(time.Now().UnixNano())

		createdDaysAgo := rand.Intn(600)
		updatedDaysAgo := createdDaysAgo / (rand.Intn(600) + 1)
		minute1 := rand.Intn(1440)
		minute2 := rand.Intn(1440)
		birthday := strfmt.Date(time.Now().Add((-time.Duration(20+rand.Intn(40))*365 + time.Duration(rand.Intn(365))) * (24 * time.Hour)))

		u := entities.User{
			ID:             int64(i),
			Gender:         "F",
			Nickname:       femaleNickname[rand.Intn(len(femaleNickname))],
			NthChild:       femaleNthChild[rand.Intn(len(femaleNthChild))],
			Birthday:       birthday,
			Tweet:          tweet[rand.Intn(len(tweet))],
			Introduction:   introduction[rand.Intn(len(introduction))],
			ResidenceState: state[rand.Intn(len(state))],
			HomeState:      state[rand.Intn(len(state))],
			Education:      education[rand.Intn(len(education))],
			Job:            job[rand.Intn(len(job))],
			AnnualIncome:   income[rand.Intn(len(income))],
			Height:         height[rand.Intn(len(height))],
			BodyBuild:      bodyBuild[rand.Intn(len(bodyBuild))],
			MaritalStatus:  maritalStatus[rand.Intn(len(maritalStatus))],
			Child:          child[rand.Intn(len(child))],
			WhenMarry:      whenMarry[rand.Intn(len(whenMarry))],
			WantChild:      wantChild[rand.Intn(len(wantChild))],
			Smoking:        smoking[rand.Intn(len(smoking))],
			Drinking:       drinking[rand.Intn(len(drinking))],
			Holiday:        holiday[rand.Intn(len(holiday))],
			HowToMeet:      howToMeet[rand.Intn(len(howToMeet))],
			CostOfDate:     costOfDate[rand.Intn(len(costOfDate))],
			Housework:      houseWork[rand.Intn(len(houseWork))],
			CreatedAt:      strfmt.DateTime(time.Now().AddDate(0, 0, -createdDaysAgo).Add(-time.Duration(minute1) * time.Minute)),
			UpdatedAt:      strfmt.DateTime(time.Now().AddDate(0, 0, -updatedDaysAgo).Add(-time.Duration(minute2) * time.Minute)),
		}
		r.Create(u)
	}
}
