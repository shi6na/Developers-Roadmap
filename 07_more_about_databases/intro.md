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

## CAP Theorem

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
