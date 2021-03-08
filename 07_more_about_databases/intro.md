# More about Databases

## ORMs

### ORM：DBを扱いやすくする「プログラミング技法」

オブジェクト関係マッピング（Object-relational Mapping）とは、データベースとオブジェクト指向プログラミング言語間の非互換なデータを変換する「**プログラミング技法**」である。

- 「インピーダンス・ミスマッチ」を解消し、RDBのレコードを直感的に扱えるようにする。
- RDBにアクセスするプログラムを書くときの煩雑な処理を軽減させ、プログラマがSQLを意識することなくコードを書ける。

#### ✔ メリット

- SQLを書かなくても良い（分かりきっていることを何回も書かなくても良い）
- データとオブジェクト（クラス）のマッピングを行わなくても良い
- `book.title #Book Title1`のように、オブジェクト指向の手法で直感的にDB操作を行えるようになる

### 「インピーダンスミスマッチ」とは？

- オブジェクト指向は「現実世界の物事に即したデータモデル」である。

- 関係データベースは「検索やCRUDなどの処理に最適化されたデータモデル」である。

このような設計思想の違いを「**インピーダンス・ミスマッチ**」と呼ぶ。

つまり、オブジェクト指向とデータベースの考えの差分を吸収して、どっちでもいい感じに使えるようにする技法がORM。

### ORMの例

- Acrive Record（Ruby on Rails）
- Eloquent（Laravel）

などが代表的なORM。

例としてAcrive Recordを使うと、以下の仕組みが得られる。

- モデルとそのデータを表す仕組み
- モデル間の関連性を表す仕組み
- 関連するモデルを通した階層の継承を表す仕組み
- DBに保存する前に検証する仕組み
- オブジェクト指向の手法でDB操作を実行する仕組み

クラスとテーブルの対応には、名前に関する規約が利用される。

- テーブル名がBookの場合、クラス名もBook
- 規約通りにクラスを作成すれば、中身が空でも正しくマッピングされる。
- 規約を外れる命令が必要な場合は、クラスに記述を追加すれば良い。

テーブルを作った後のActive Recordの動きは以下。

- ActiveRecordライブラリがDBのスキーマを実行時に読み取る
- カラム名と同じ名前の属性を使えるように、カラムのデータ型に従って、属性が適切なRubyのクラスへの対応付けが行われる。

<!-- 便利だな🎶と思って使っていたが、これが無かったら手動でマッピングしないとオブジェクト操作が出来なかったとは… -->

## Transactions

ソフトウェアの処理方式の1つ。互いに関連・依存する複数の処理をまとめ、**分割できない1つの処理単位**として扱うこと。

- 銀行の口座振込処理など、成功か失敗のどちらかしか許容されない処理を実装する時に使われる。
- 「振り込めた」か「振り込めてない」かの2択しかない。「ちょっとだけ振り込めた」とか無い。

### SQL（PostgreSQL）におけるトランザクションの実装方法

PostgreSQLでは、トランザクションを構成するSQLコマンドを**BEGIN**と**COMMIT**で囲んで設定する。

```zsh
BEGIN;
  構成したいSQL;
COMMIT;
```

例えば、銀行のデータベースで、そこに多数の顧客の口座残高と支店の総預金残高が含まれているとする。

その時、アリスの口座からボブの口座に$100.00の支払いがあったことを記録したい場合、以下のようなSQLになる。

```zsh
UPDATE accounts SET balance = balance - 100.00
    WHERE name = 'Alice';
UPDATE branches SET balance = balance - 100.00
    WHERE name = (SELECT branch_name FROM accounts WHERE name = 'Alice');
UPDATE accounts SET balance = balance + 100.00
    WHERE name = 'Bob';
UPDATE branches SET balance = balance + 100.00
    WHERE name = (SELECT branch_name FROM accounts WHERE name = 'Bob');

```

- これをBEGINとCOMMITで囲んで実行すると、トランザクションが構成される。
- トランザクションの途中で、コミットを行わないと判断（口座残高が足りない場合など）した際は、ROLLBACKを使用して、行われた全ての更新を破棄するようになっている。

A. トランザクションが成功した場合

```zsh
mydb=# BEGIN;
BEGIN
mydb=*# UPDATE accounts SET balance = balance - 100.00
    WHERE name = 'Alice';
UPDATE branches SET balance = balance - 100.00
    WHERE name = (SELECT branche_name FROM accounts WHERE name = 'Alice');
UPDATE accounts SET balance = balance + 100.00
    WHERE name = 'Bob';
UPDATE branches SET balance = balance + 100.00
    WHERE name = (SELECT branche_name FROM accounts WHERE name = 'Bob');
UPDATE 1
UPDATE 1
UPDATE 1
UPDATE 1
mydb=*# COMMIT;
COMMIT
mydb=# SELECT * FROM accounts;
 name  | balance | branche_name
-------+---------+---------------
 Alice |     100 | Hayward
 Bob   |     300 | San Francisco
(2 rows)
```

B. 失敗した場合（Aliceの口座残高が不足していた場合）

```zsh
mydb=# BEGIN;
BEGIN
mydb=*# UPDATE accounts SET balance = balance - 100.00
    WHERE name = 'Alice';
UPDATE branches SET balance = balance - 100.00
    WHERE name = (SELECT branche_name FROM accounts WHERE name = 'Alice');
UPDATE accounts SET balance = balance + 100.00
    WHERE name = 'Bob';
UPDATE branches SET balance = balance + 100.00
    WHERE name = (SELECT branche_name FROM accounts WHERE name = 'Bob');
ERROR:  new row for relation "accounts" violates check constraint "accounts_balance_check"
DETAIL:  Failing row contains (Alice, 0, Hayward).
ERROR:  current transaction is aborted, commands ignored until end of transaction block
ERROR:  current transaction is aborted, commands ignored until end of transaction block
ERROR:  current transaction is aborted, commands ignored until end of transaction block
mydb=!# COMMIT;
ROLLBACK
```

#### 余談

- PostgreSQLは実際、すべてのSQL命令文をトランザクション内で実行するようになっている。
- BEGINを書かずとも、それぞれの命令文は暗黙的なBEGINがついているとみなし、成功すればCOMMITで囲まれたものとしている。

### FW（Ruby on Rails）におけるトランザクションの実装方法

```ruby
ActiveRecord::Base.transaction do
  例外が発生するかもしれない処理
end
  例外が発生しなかった場合の処理
resque => e
  例外が発生した場合の処理
```

```ruby
モデル.transaction do
  例外が発生するかもしれない処理
end
  例外が発生しなかった場合の処理
resque => e
  例外が発生した場合の処理
```

railsにおいては、上記の方法でトランザクションを実装することが出来る。

#### ✔ 特徴

- ブロック内の全ての処理が正常に行われた場合に保存が行われる
- エラーが発生した場合は、ロードバックをする
- 複数のデータベースにまたがる分散トランザクションはサポートしていない
- 使用するにはDBがトランザクションをサポートしていることが必要

以下、rails consoleでUserモデルを用いて実際にトランザクション処理を行ってみた結果。

A. 処理成功する場合🎉

```irb
irb(main):001:1> User.transaction do
irb(main):002:1*   a1 = User.new(name: 'tarou', email: 'tarou@email.com')
irb(main):003:1>   a1.save!
irb(main):004:1>   a2 = User.new(name: 'jurou', email: 'jirou@email.com')
irb(main):005:1>   a2.save!
irb(main):006:1> end
   (0.7ms)  SELECT sqlite_version(*)
   (0.0ms)  begin transaction
  User Create (0.7ms)  INSERT INTO "users" ("name", "email", "created_at", "updated_at") VALUES (?, ?, ?, ?)  [["name", "tarou"], ["email", "tarou@email.com"], ["created_at", "2021-03-01 06:09:48.910751"], ["updated_at", "2021-03-01 06:09:48.910751"]]
  User Create (0.1ms)  INSERT INTO "users" ("name", "email", "created_at", "updated_at") VALUES (?, ?, ?, ?)  [["name", "jurou"], ["email", "jirou@email.com"], ["created_at", "2021-03-01 06:09:48.917575"], ["updated_at", "2021-03-01 06:09:48.917575"]]
   (0.5ms)  commit transaction
=> true
```

B. 失敗する場合（emailが未入力）

```irb
irb(main):007:0> User.transaction do
irb(main):008:1*   a1 = User.new(name: 'tarou', email: 'tarou@email.com')
irb(main):009:1>   a1.save!
irb(main):010:1>   a2 = User.new(name: 'jurou')
irb(main):011:1>   a2.save!
irb(main):012:1> end
   (0.1ms)  begin transaction
  User Create (0.4ms)  INSERT INTO "users" ("name", "email", "created_at", "updated_at") VALUES (?, ?, ?, ?)  [["name", "tarou"], ["email", "tarou@email.com"], ["created_at", "2021-03-01 06:10:35.460856"], ["updated_at", "2021-03-01 06:10:35.460856"]]
   (0.3ms)  rollback transaction
Traceback (most recent call last):
        2: from (irb):7
        1: from (irb):11:in `block in irb_binding'
ActiveRecord::RecordInvalid (Validation failed: Email can't be blank)
```

### ロールバックて何よ

- データ更新などで障害が起こった時に、**その前の状態にまで戻ること**を言う。後進復帰とも。
- 障害発生時に、処理途中で確約されていないトランザクションの処理をすべて取り消し、最初の時点に戻す。

## ACID

信頼性のあるトランザクションシステムの持つべき性質として定義された概念。以下4つの性質の頭文字を取ったもの。

**A**tomicity：不可分性

**C**onsistency：一貫性

**I**solation：独立性

**D**urability：永続性

### **A**tomicity：不可分性

トランザクションに含まれるタスクが全て実行されるか、あるいは全く実行されないことを保証する性質を言う。

「**原子性**」とも呼ばれる。

#### 例

口座AB間で送金が行われる際、送金操作は以下の2操作によって行われる。

- 口座Aの残高から1万円引く
- 口座Bの残高に1万円加える

「不可分性が保証される」とは、上記2操作が**全て行われる**か、あるいは**全く行われない**ことを指す。

### **C**onsistency：一貫性

トランザクション開始と終了時にあらかじめ与えられた整合性を満たすことを保証する性質を指す。

→ つまり、条件や整合性を満たさない状態になるようなトランザクションは実行が中断される。

#### 例

- 預金残高は、一般的に0または正の値を取る条件を満たす必要がある。
- 口座Aから送金を行うとき、その前後でAの口座残高が負になるような額は送金できないようにする。

### **I**solation：独立性

- 処理の過程が、他の操作から隠蔽されること。
- 処理の結果だけが他から見ることができ、実行中の途中状態が他へ影響することが無いということ。

#### 例

- 残高100万円の口座Aから、残高200万円の口座Bに1万円送金する場合の操作が以下の順序で行われたとする。

1. 口座Aの残高から1万円を引く
1. 口座Bの残高に1万円を加える

この時、内部状態として取り得るのは以下の３つになる。

| 時点 | 口座A | 口座B |
| --- | --- | --- |
| 送金前 | 100万円 | 200万円 |
| 実行中 | 99万円 | 200万円 |
| 送金後 | 99万円 | 201万円 |

独立性が保証されている場合は、外部からは**送金前**と**送金後**のいずれかの状態しか観測できない。

### **D**urability：永続性（耐久性）

処理の完了通知をユーザが受けた時点で、その操作は永続的なものとして、データベースに記録されていることを指す。

- ≒「システム障害に耐える」。
- 処理を永続性記憶装置（不揮発性メモリ。HDD、CDなど）上にログとして記録し、システムに異常が発生した場合、そのログを用いて異常発生前の状態まで復旧する。

### それぞれの性質を実現させる代表的な機能

| 性質 | 機能 | 概要 |
| --- | --- | --- |
| 原子性 | コミットメント制御 | コミット機能とロールバック機能のこと。コミット機能とは、トランザクション処理が全て実行された時、その**処理結果を確定させ、データを更新する**機能。|
| 一貫性・独立性 | 排他制御 | 共有資源に対しての複数同時アクセスにより不整合が発生することを防ぐため、あるトランザクションが共有資源にアクセスしている時は、他トランザクションからはアクセスできないようにして直列に処理されるように制御すること。 |
| 永続性 | 障害回復機能 | バックアップやリストア機能など、障害発生時に、障害発生前まで回復させる機能のこと。 |

## CAP Theorem：CAP定理

分散コンピュータシステムのコンピュータ間の情報複製において、

- 一貫性：**C**onsistency
- 可用性：**Av**ailability
- ネットワーク分断耐性：**P**artition-tolerance

この3つの保証のうち、同時に満たせるのは2つまでであり、同時に全てを満たすことは出来ない、という定理。ブリュワーの定理とも。

保証はそれぞれ以下を意味している。

1. 一貫性：誰かがデータを更新したら必ず更新後のデータを参照できること
1. 可用性：クライアントは必ずデータにアクセスできること
1. ネットワーク分断耐性：データを複数のサーバに分散保管できること

２つの保証を取った時、除かれる要素に応じて「CA」「CP」「AP」の3種類に分類される。

- CA：単一サーバで動作するDB
- CP：分散データベース
- AP：DNS、NTP、HTTPキャッシュ

などのシステムに用いられることが多い。3種の中ではAPが最も障害に強い。

例えば、「その単一箇所が働かないと、システム全体が障害となような箇所（単一故障点）」があった時のCAP定理での考え方は以下になる。
> ネットワーク分断が発生した際にシステムがバラバラに分裂しても、単一故障点を基準に一貫した応答が出来る（分断耐性＋一貫性）が、可用性が成立しなくなる。

## Data replication

## Sharding Strategies

## N + 1 Probrem

## Database Nomalization

## Index and how they work

## 参考文献

- [オブジェクト関係マッピング - Qiita](https://qiita.com/yk-nakamura/items/acd071f16cda844579b9)
- [オブジェクト関係マッピング - Wikipedia](https://ja.wikipedia.org/wiki/%E3%82%AA%E3%83%96%E3%82%B8%E3%82%A7%E3%82%AF%E3%83%88%E9%96%A2%E4%BF%82%E3%83%9E%E3%83%83%E3%83%94%E3%83%B3%E3%82%B0)
- [O/Rマッパーを使う理由 - ぺい](https://tikasan.hatenablog.com/entry/2018/08/06/110000)
- [transaction | Railsドキュメント](https://railsdoc.com/page/transaction)
- [トランザクションとは - IT用語辞典 e-Words](https://e-words.jp/w/%E3%83%88%E3%83%A9%E3%83%B3%E3%82%B6%E3%82%AF%E3%82%B7%E3%83%A7%E3%83%B3.html)
- [トランザクション](https://www.postgresql.jp/document/9.4/html/tutorial-transactions.html)
- [ACID特性とは何？ Weblio辞書](https://www.weblio.jp/content/ACID%E7%89%B9%E6%80%A7)
- [排他制御（楽観ロック・悲観ロック）の基礎　 - Qiita](https://qiita.com/NagaokaKenichi/items/73040df85b7bd4e9ecfc)
- [原子性を実現する「コミットメント制御」を理解する：「データベーススペシャリスト試験」戦略的学習のススメ（22） - ＠IT](https://www.atmarkit.co.jp/ait/articles/1703/01/news195.html)
- [耐久性を確保する「障害回復機能」を理解する：「データベーススペシャリスト試験」戦略的学習のススメ（25） - ＠IT](https://www.atmarkit.co.jp/ait/articles/1703/01/news198.html)
- [ACID (コンピュータ科学) - Wikipedia](https://ja.wikipedia.org/wiki/ACID_(%E3%82%B3%E3%83%B3%E3%83%94%E3%83%A5%E3%83%BC%E3%82%BF%E7%A7%91%E5%AD%A6))