package repository

var autharray = map[string]string{
	"taro": "japan",
	"mike": "america",
	"leon": "german",
	"次郎":   "日本",
}

// 同様のデータをmySQLに入れる場合以下のクエリを実行
/*
INSERT INTO users (name, pass) VALUES('taro', '12345');
INSERT INTO users (name, pass) VALUES('mike', 'abcde');
INSERT INTO users (name, pass) VALUES('leon', 'ABCDE');
INSERT INTO users (name, pass) VALUES('次郎', '123Ab');
*/
