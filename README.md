# board-games
## 概要
所謂「ボードゲームアリーナ」的なもの。  
プロジェクトが軌道にのれば各々が好きなゲームを作っていくスタイルにしたい。  

## 開発方針
まず初めに"アカウント登録機能"を用意する。  
次から早速ボードゲームのAPI実装に入るが、初めは簡単なものから作っていく。  
e.g.) ヌメロン, リバーシ, ババ抜き...  
作る際はissueにチケットを作成して「誰が何を」作成しているのか可視化出来るようにする。  
ベースbranchは main->feature/game-title とする。  
1つのゲームに複数の機能が存在するので、作業branchは feature/game-title->feature/game-title/function1, feature/game-title/function2 とする。  
PullRequest（以降PR）は出来るだけ機能単位で作成する。  
アーキテクチャ選定はDDD+レイヤードアーキテクチャとする。  
（参考: [今すぐ「レイヤードアーキテクチャ+DDD」を理解しよう](https://qiita.com/tono-maron/items/345c433b86f74d314c8d) ）  
**全員初学者なので羞恥心や遠慮は捨てること**  
**また質問やクソコードを馬鹿にしないこと**

## 技術選定
backendはGoを使用。  
frontendはなんでもいいが, Jsになると思う。  
gopherjsなるものが存在し、Goでフロント実装も出来るが最初のうちはそこまでしなくて良い。  
DataBaseはMySQLを使用。  
理由は
1. ゲームデータの特徴としてユーザをprimary keyとしたレコードが多いため。
1. ボードゲームのレコードは大半がログでありRDBMSが好ましいため。

（参考: [ゲームエンジニアのためのデータベース設計](https://www.slideshare.net/sairoutine/ss-62485460) ）

## 設計
初めは、browser <-> local API -> local storage  
形になってくればAWSに置いて公開したい。

## 開発する前に見ておいた方が良いもの
1. [Go公式サイト](https://golang.org/)
1. [Go言語基礎文法最速マスター](https://go-tour-jp.appspot.com/welcome/1)
1. [Goの基礎](https://qiita.com/tfrcm/items/e2a3d7ce7ab8868e37f7)
1. メルカリやDeNAの記事（日本では先駆者）