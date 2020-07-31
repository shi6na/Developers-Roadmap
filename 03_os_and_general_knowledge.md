# OS and General Knowledge

## Terminal Usage

- `command + n`で新規ウィンドウ
- `command + t`で新規タブ

### ユーザモード

- 一般的なUNIX系のシステムでは、ユーザはスーパーユーザと一般ユーザに分けられる。  
- スーパーユーザはあらゆる操作が可能で、ユーザー名はrootに決められている。  
- macOSでは、デフォルトでrootにログインしたり、一時的にrootに移行して操作するといったことが出来ないようになっているため、**sudo**というコマンドを利用して、一時的にスーパーユーザの権限を取得し、コマンドを実行している。

### SIP(System Integrity Protection)

- sudoコマンドを使えばシステムに対してあらゆる操作が行えるため、システムに重大な影響を与えてしまう可能性がある。  
- そのため、SIPと呼ばれる仕組みで、スーパーユーザであってもシステムの重要なディレクトリ内の書き換えが出来ないようになっている。  
例えば、下記のディレクトリ以下はsudoコマンドを使っても変更することができないようになっている。（表示は可）

- `/System`
- `/bin`
- `/sbin`
- `/user`

これらを書き換えたい時はリカバリーモードで起動して、`csrutil`コマンドを使う。  
ターミナルのデザインや書式設定は環境設定からどうぞ。  

### プロンプトの変更

プロンプトの書式は、デフォルトでは「`ホスト名：現在のディレクトリ名 ユーザ名$`」になっているが、  
必要に応じてこの書式を変更することが出来る。  

プロンプトの書式は、シェル変数「PS1」で定義されている。  
例えば、`ホスト名：現在のディレクトリ名 ユーザ名$` 👉 `\h:\W \u\$` といった書式で設定されている。  
書式で使えるエスケープ文字は以下の通り。  

| エスケープ文字 | 意味 |
| --- | --- |
| `\H` | ホスト名 |
| `\h` | ホスト名(現在の.まで) |
| `\W` | 現在のディレクトリ名 |
| `\w` | 現在のディレクトリ（フルパス） |
| `\u` | 現在のユーザー名 |
| `\$` | $は変数を表すので、$を表示させたい場合は\$のようにエスケープ |
| `\T` | 時刻 HH:MM:SS 形式 (12時間) |
| `\t` | 時刻 HH:MM:SS 形式 (24時間) |

一時的にプロンプトを変更するには、以下のようにPS1に書式を指定する。  
`PS1="\W \$ "  #現在のディレクトリ名 $`  
上記の場合、ターミナルを再起動すると設定が戻ってしまうので、設定を保存するには~/bashrcなどに設定を記述する。  

### コマンド履歴

- 実行したコマンドはコマンド履歴として保存される。  
- `history [表示するコマンド数]`で過去に実行したコマンドの一覧を表示することが出来る。引数なしだと最大500個。  
- 表示したコマンドの中から、`![コマンド番号]`のように、番号を指定してコマンドを実行することができる。  
- `!!`と実行すると、直前に実行したコマンドを再実行することができる。  

#### インクリメンタルサーチ

- コマンド履歴の中から特定のコマンドを検索するには、`control + r`を押してインクリメンタルサーチというモードに切り替える。  
- `` (reverse-i-search)`': `` と表示され、続けて検索するコマンド名を入力すると、新しい方から古い方に向かって入力したコマンドが検索され、表示される。  
- `control + r`を押し続けていると、同じ文字列を含む次に古いコマンドが検索される。  
- `esc`でコマンドがプロンプトの後に入力された状態になるので、そこから実行が可能。  

#### マーク機能

過去に実行したコマンドの結果などを確認したい時に、マークやブックマーク機能が便利。  
それぞれ設定されているマークやブックマークの位置に移動することが出来る。  

- 角括弧`[]`：マーク
- 縦線`|`：ブックマーク

ターミナルの設定で自動的に各プロンプトの行を自動的にマークする他、  
マークしたい行を選択して、ショートカットメニューから「マーク」(`command + u`)を選択。  
マークを削除したいときには、ショートカットメニューから「マークを解除」(`shift + command + u`)を選択。  

ブックマークは、マークに名前を設定出来るようにしたもので、動作としても重みのある「マーク」。  
ブックマークを挿入するには、ターミナルのメニューから「ブックマークを挿入」(`shift + command + u`)を選択。  

マークやブックマークの位置に移動するには、ターミナルメニューの「編集」👉「移動」から選択する他、  
以下のショートカットキーが使用できる。  

- 前のマークへジャンプ(`command + ↑`)
- 次のマークへジャンプ(`command + ↓`)
- 前のブックマークへジャンプ(`option + command + ↑`)
- 次のブックマークへジャンプ(`option + command + ↓`)

### ディレクトリを表す記号

| 記号 | 意味 |
| --- | --- |
| ~(チルダ) | ホームディレクトリ |
| ~-(チルダとハイフン) | 直前にいたディレクトリ |
| .(ドット) | カレントディレクトリ |
| ..(ドット２つ) | 1つ上のディレクトリ |

### ワイルドカード

| 記号 | 意味 |
| --- | --- |
| * | 任意の文字列（長さ0以上の文字列）とマッチ |
| ? | 任意の1文字とマッチ |
| [文字] | 指定した文字（複数や範囲を指定可能）の中の１文字とマッチ |
| [^文字]または[!文字] | 指定した文字（複数や範囲を指定可能）の中以外の１文字とマッチ |

### コマンドを連続して実行

#### 実行結果に関わらず、コマンド１が終了したらコマンド2を実行( `;` )

`command1 ; command2`
> command1が失敗しても、command2は実行されます`

#### コマンド1が成功したらコマンド2を実行する( `&&` )

`command1 && command2`
> command1が成功した場合のみ、command2が実行されます。

※前のコマンドが失敗すると、次のコマンドは実行されません。

#### コマンド1が失敗した場合のみコマンド2を実行する( `||` )

`command1 || command2`
> command1が失敗した場合のみ、command2が実行されます

### UNIX系システムの主なディレクトリ

| ディレクトリ | 概要（配置されるファイル） |
| --- | --- |
| /bin、/usr/bin | 基本コマンド |
| /sbin、/usr/sbin | システム管理用コマンドなど |
| /dev | デバイスファイル |
| /etc | 設定ファイル。/private/etcのシンボリックリンク。 |
| /tmp | 一時的なファイル。ここに配置されたファイルは次回のシステム起動時には削除される。/private/tmpのシンボリックリンク。 |
| /var | 変更されるデータ。キャッシュデータ、ログファイルなど変更されるファイル。/private/tmpのシンボリックリンク。 |
| /usr/lib | ライブラリファイル |
| /usr/local | ユーザがインストールしたソフトウェアのファイル |
| /usr/share | システムに依存しない共有データファイルやドキュメント、オンラインマニュアルなど |

### macOC独自の主なディレクトリ

| ディレクトリ | 概要（配置されるファイル） |
| --- | --- |
| /Users | ユーザのホームディレクトリ。Finderでは「ユーザ」。 |
| /System/Library | macOS独自の拡張機能などのライブラリファイル。Finderでは「システム」👉「ライブラリ」 |
| /Library | アプリケーション独自の拡張機能などのライブラリファイル。Finderでは「ライブラリ」 |
| /Applications | アプリケーションのファイル。Finderでは「アプリケーション」|

<!-- 余力があればリダイレクトのことについて -->

## How OSs work in General

## OSとは

- オペレーティングシステム(Operating System)
- OSはアプリケーションシステム（Excel、ウェブブラウザなど）を動作させることが目的
- OSはあらゆるハードとソフトウェアを管理し、利用者が利用しやすいようにサービスしてくれるもの

## OSの構成

### ブートローダー

- コンピュータを起動したときに呼び出される、OSをロードするプログラム
- ブート（起動）ローダー（読み込む）
- 電源投入 👉 BIOS起動 👉 **ブートローダー起動** 👉 OS起動

### カーネル

- カーネル = 核
- OSの中核部分のソフトウェア
- ハードウェアの操作を一手に引き受けるソフトウェア
- プロセス管理、空間管理（メモリのユーザー空間など）、時間管理、割り込み処理、ファイルシステム、ネットワークなどの役割を担っている

### デーモン

- OSが起動すると同時に動く**常駐プログラム**
- 例えばWebサーバー(httpd)として機能していたり、FTPサーバー(ftpd)として起動していたりする
- 上記のように、デーモンのプログラム名には`d`が付いていることが多い。

### シェル

- OSのユーザのためにインタフェースを提供するソフトフェアであり、カーネルのサービスへのアクセスを提供する。
- シェル = 外殻。カーネルとユーザの間にある外殻であることから。
- CLIのものとGUIのものがあり、CLIは操作が早かったりする一方、GUIはユーザビリティに優れ、画像や動画の操作に適している。
- Macに於いて、主なCLI（コマンドラインシェル）は「`zsh`」、GUI（グラフィティカルシェル）は「`Finder`」。

### デスクトップマネージャ

- デスクトップ環境を提供するソフトウェアのこと。  
- WIMP、ツールバー、フォルダ、背景画像、デスクトップウィジェットなどといったものから成り立っている。

※WIMP 👉 Window,Icon,Menu,Pointer

### アプリケーション

- OSだけの機能ではユーザビリティが低いため、 OSには一般のユーザーが使いやすいように様々なソフトウェアがセットになっている。
- 例えば、ファイル管理、ブラウザ、各種設定、メールなどの標準アプリケーションなど。

## Process Management

- プロセスの生成・実行・消滅を管理すること。
- OSのうちでも主に**カーネル**の機能。
- プロセス間通信や排他制御もプロセスマネジメントの役割。
- プロセスへのリソースを割り当てる機構でもある。

- プロセスとは：命令。動作中のプログラム。ソースコードや、それを実行するためのリソースも含まれる。

## Threads and Concurrency

### スレッドとは

- 処理の単位。実行の文脈とも。実行単位はプロセスよりも小さい。
- プロセスに比べて、プログラムを実行する時のコンテキスト情報が最小で済むため切り替えが早い。
- プロセスは親プロセスから子プロセスが作られ、木構造になる上に、子プロセスに仮想メモリが与えられるので、重い処理になる。
- そこで、メモリなど親子で共有できるところは共有したのが「スレッド」という実行単位。 👉 だからコンテキスト情報が最小で済む
- プロセスが複数のスレッドで構成されている場合もあり（マルチスレッド）、この場合は命令を同時に実行する。👉 並列処理や並行処理

### 並行処理とは

- ある1つの時点では1つの仕事しかしていないが、複数の仕事を切り替えることによって、同時にやっているように見せる処理方式のこと。
- 「他を待たせないこと（待ち時間を上手く使うこと）」「同時にやること」が目的。
- 実際に物理的に複数の処理を同時に実行しているわけではない（それは並列処理）。1つのCPUで高速に実行タスクを切り替えて、同時にやっているように見せているだけ。

### 並列処理とは

- 実際に物理的に複数の処理を同時に実行する処理方式のこと。
- 複数のCPUを用いて実行する。ただし、マルチコアプロセッサに対応したシステムでないと実現できないので注意。
- 単一の処理装置を用いる場合に比べ、最大で装置の数を乗じた処理性能（プロセッサ4基なら4倍）を発揮できる可能性がある。

## Basic Terminal Commands

- `$`はbash、`%`はzsh。

### grep

`grep [オプション] 検索パターン ファイル`

- ファイル内検索コマンド。
- ファイルを引数で指定しなかった場合は、標準入力から読み込む。
- オプションにより正規表現も使える。

#### 実行してみた - 普通にgrep

```zsh
% grep OS 03_os_and_general_knowledge.md
# OS and General Knowledge
macOSでは、デフォルトでrootにログインしたり、一時的にrootに移行して操作するといったことが出来ないようになっているため、  
| /System/Library | macOS独自の拡張機能などのライブラリファイル。Finderでは「システム」👉「ライブラリ」 |
## How OSs work in General
## OSとは
- OSはアプリケーションシステム（Excel、ウェブブラウザなど）を動作させることが目的
- OSはあらゆるハードとソフトウェアを管理し、利用者が利用しやすいようにサービスしてくれるもの
## OSの構成
- コンピュータを起動したときに呼び出される、OSをロードするプログラム
- 電源投入 👉 BIOS起動 👉 **ブートローダー起動** 👉 OS起動
- OSの中核部分のソフトウェア
- OSが起動すると同時に動く**常駐プログラム**
- OSのユーザのためにインタフェースを提供するソフトフェアであり、カーネルのサービスへのアクセスを提供する。
- OSだけの機能ではユーザビリティが低いため、 OSには一般のユーザーが使いやすいように様々なソフトウェアがセットになっている。
- OSのうちでも主に**カーネル**の機能。
## POSIX Basics
```

- `OS`を片っ端から検出してくれている。POSIXまで引っかかってる。

#### 実行してみた - OR検索

```zsh
% grep "bash\|zsh" 03_os_and_general_knowledge.md
上記の場合、ターミナルを再起動すると設定が戻ってしまうので、設定を保存するには~/bashrcなどに設定を記述する。  
- Macに於いて、主なCLI（コマンドラインシェル）は「`zsh`」、GUI（グラフィティカルシェル）は「`Finder`」。
- `$`はbash、`%`はzsh。
```zsh
```zsh
```zsh
```

- `bash`と`zsh`の複数パターン（OR）検索をしてくれている。
- `-E` を付けると明示的に検索パターンを正規表現として読んでくれるため、バックスラッシュが要らない。

### awk

`awk 'スクリプト' 入力ファイルのパス`

- 「オーク」と読む。コマンドラインから簡単にテキストをフィルターしたり、表示を整えたり、値を集計したりできる。
- ポイントは「コマンドラインから簡単にテキストを操作する」こと。コマンドラインから出力したテキストをその場で処理するのに大変便利。
- コマンドっぽく使えるだけで、正確にはコマンドではなく「AWKスクリプト・インタプリタ。」
- スクリプトの箇所は「パターン文」と「アクション文」から成り立つ。
- `awk 'パターン文 {アクション文}' 入力ファイルのパス`

#### 組込変数

| 名称 | 説明 | デフォルト値 |
| --- | --- | --- |
| $0 | レコード全体 |
| $n | レコード(列)のn番目 |
| RS | Record Separator - 入力のレコード区切り文字 | 改行 |
| FS | Field Separator - 入力のフィールド区切り文字 | 連続するスペースorタブ文字 |
| ORS | Output Record Separator - 出力のレコード区切り文字 | 改行 |
| OFS | Output Field Separator - 出力のフィールド区切り文字 | スペース１つ |
| NR | Number of Record - 現在のレコード数 |
| NF | Number of Field - 現在のレコードのフィールド数 |

#### 実行してみた - 入力ファイルの中身を出力

```zsh
% awk '{print NR "\t" $0}' 02_basic_frontend_knowledge/sample.html
1       <!DOCTYPE html>
2       <html lang="ja">
3       <head>
4         <meta charset="UTF-8">
5         <meta name="viewport" content="width=device-width, initial-scale=1.0">
6         <link rel="stylesheet" href="style.css">
7         <link href="https://fonts.googleapis.com/css2?family=Kosugi+Maru&family=M+PLUS+Rounded+1c:wght@300;800&display=swap" rel="stylesheet">
8         <title>Developer Roadmap</title>
9       </head>
10      <body>
11        <div class="main">
12          <div class="card">
13            <h1 class="title">フォームに入力した文字を下に表示するやつ</h1>
14            <form action="#">
15              <input type="text" id="input_message" class="textbox">
16              <input type="button" class="btn_submit" value="送信" onclick="getInput()">
17            </form>
18            <p id="output_message" class="output_message"></p>
19            <script src="getInput.js"></script>
20          </div>
21        </div>
22      </body>
23      </html>
```

- `NR "\t"`で行番号を表示させることが可能。

#### 実行してみた - パターン一致

```zsh
% awk '/meta/ {print $0}' 02_basic_frontend_knowledge/sample.html
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
```

- パターンは正規表現で。

#### 実行してみた - 集計

例えば、rootが起動しているプロセスのCPU使用率合計が知りたい時。  
まずは全てのプロセスとCPU使用率を表示。

```zsh
% ps aux
USER               PID  %CPU %MEM      VSZ    RSS   TT  STAT STARTED      TIME COMMAND
miyasato-pc       1461  84.1  3.8  9111452 1266244   ??  S     9:32AM 325:54.66 /Applications/Snap Camera.app/Contents/MacOS/Snap Camera
miyasato-pc       3848  16.5  0.7 25628312 242376   ??  S     2:00PM   3:35.13 /Applications/Google Chrome.app/Contents/Frameworks/Google Chrome Framework.framework/Versions/84.0.4147.89/Helpers/Google Chr
〜〜〜
〜〜〜
miyasato-pc       3896   0.0  0.0  4341312   2468   ??  S     2:05PM   0:00.01 /System/Library/PrivateFrameworks/ToneLibrary.framework/Versions/A/XPCServices/com.apple.tonelibraryd.xpc/Contents/MacOS/com.a
```

この出力を保存してrubyなり何なりで手を加えても良いが、awkを使うと全表示も含め1行で済む。

```zsh
% ps aux | awk '$1 == "root" {s += $3} END {print s}'
0.4
```

平和。

### sed

`sed [オプション] スクリプトコマンド 入力ファイル`

- **Stream EDitor**の略で、指定したファイルをコマンドに従って処理し、標準出力へ出力する。
- コマンドと対象となる「スクリプト」によって、テキストファイルを編集するコマンド。テキストファイルに直接手を加えるわけではない。(-iオプションをつけたら直接編集して上書き)
- 例えば、「`sed s/abc/ABC/ ファイル名`」で、指定したファイル内の「`abc`」を「`ABC`」に置き換えることができるが、この「`s/abc/ABC/`」部分が、sedのスクリプト。
- `s`は「置換する」というコマンドで、それ以下が正規表現。区切り記号の「` / `」は他の記号でもよく、パスの置換などで置換対象に「` / `」が含まれている場合は、「` ! `」など他の記号を使ったほうが便利。
- パイプラインやリダイレクトを活用するのが一般的。

#### 実行してみた - binをBINに置き換えて出力

```zsh
% cat /etc/shells | sed s/bin/BIN/
# List of acceptable shells for chpass(1).
# Ftpd will not allow users to connect who are not using
# one of these shells.

/BIN/bash
/BIN/csh
/BIN/dash
/BIN/ksh
/BIN/sh
/BIN/tcsh
/BIN/zsh
```

- `sed s/bin/BIN/ /etc/shells`でも同様の結果。

### lsof

- 「**LiSt Open Files**」（開いているファイル群を列挙する）
- 「プロセスが開いているファイル」を表示するコマンド。
- >UNIX／Linuxでは、一般的なファイルだけでなく、ネットワークソケットやデバイスドライバー、プロセス情報なども「ファイル」として扱います。そのため、開かれているファイルを調べることで、待機ポートやネットワークのマッピング情報などを把握できます。
- >lsofコマンドによって「あるファイルを開いているプロセス」「あるポートにアクセスしているプロセス」を特定し、「不要なプログラムが実行されていないか」「不正なプログラムが動作していないか」を調べることができます。

#### 実行してみた - TCP通信のプロセスを調べる

```zsh
% lsof -iTCP
COMMAND     PID        USER   FD   TYPE            DEVICE SIZE/OFF NODE NAME
Notion      507 miyasato-pc   47u  IPv4 0x8e0459f36a5232d      0t0  TCP 10.10.111.55:64149->104.18.23.110:https (ESTABLISHED)
Notion      507 miyasato-pc   73u  IPv4 0x8e0459f3696294d      0t0  TCP 10.10.111.55:62415->ec2-34-237-73-95.compute-1.amazonaws.com:https (ESTABLISHED)
Slack\x20   568 miyasato-pc   20u  IPv4 0x8e0459f295bd58d      0t0  TCP 10.10.111.55:61889->ec2-18-178-165-242.ap-northeast-1.compute.amazonaws.com:https (ESTABLISHED)
~~~
~~~
Snap\x20C  1461 miyasato-pc   35u  IPv4 0x8e0459f3ac431cd      0t0  TCP 10.10.111.55:64501->nrt13s50-in-f19.1e100.net:https (ESTABLISHED)
git-crede  1884 miyasato-pc    5u  IPv4 0x8e0459f2ec791cd      0t0  TCP 10.10.111.64:50121->ec2-13-114-40-48.ap-northeast-1.compute.amazonaws.com:https (CLOSED)
com.docke  9677 miyasato-pc   25u  IPv6 0x8e0459f2bdd7c9d      0t0  TCP *:5506 (LISTEN)
com.docke  9677 miyasato-pc   26u  IPv6 0x8e0459f2bddad9d      0t0  TCP *:hbci (LISTEN)
com.docke 10342 miyasato-pc   10u  IPv4 0x8e0459f295bf32d      0t0  TCP localhost:57381 (LISTEN)
```

#### 実行してみた - ポート番号80(`http`)を使用しているプロセスを調べる

```zsh
% lsof -i:80
COMMAND    PID        USER   FD   TYPE            DEVICE SIZE/OFF NODE NAME
Google    1081 miyasato-pc   30u  IPv4 0x8e0459f2ec7af6d      0t0  TCP 10.10.111.58:56418->nrt12s23-in-f14.1e100.net:http (CLOSE_WAIT)
Google    1081 miyasato-pc   40u  IPv4 0x8e0459f36da232d      0t0  TCP 10.10.111.58:56419->p077.net027121054.biz.tokai.or.jp:http (CLOSE_WAIT)
```

- `-P`をつけるとサービス名(`http`)に変換されずポート番号(`80`)のまま表示される。

```zsh
% lsof -i:80 -P
COMMAND    PID        USER   FD   TYPE            DEVICE SIZE/OFF NODE NAME
Google    1081 miyasato-pc   30u  IPv4 0x8e0459f2ec7af6d      0t0  TCP 10.10.111.58:56418->nrt12s23-in-f14.1e100.net:80 (CLOSE_WAIT)
Google    1081 miyasato-pc   40u  IPv4 0x8e0459f36da232d      0t0  TCP 10.10.111.58:56419->p077.net027121054.biz.tokai.or.jp:80 (CLOSE_WAIT)
```

#### 補足 - TCPの状態

| 状態 | 説明 |
| --- | --- |
| LISTEN | 接続待受状態|
| ESTABLISHED | 接続が確立されている状態 |
| SYN_SENT | 接続要求(SYN)を送信した状態(応答/ACKは受けてない状態) |
| SYN_RECV | 接続要求(SYN)を受け取った状態 |
| FIN_WAIT1 | ソケットを閉じ、接続を落としている状態 |
| FIN_WAIT2 | 接続はクローズされ、ソケ ットはリモート側からの切断を待っている状態 |
| TIME_WAIT | 接続終了を待っている状態 |
| CLOSED | ソケットが使用されていない |
| CLOSE_WAIT | 接続相手はクローズし、自身はクローズ待ちの状態 |
| LAST_ACK | FINに対する応答(ACK)待ち |
| CLOSING | FIN_WAIT1でFINを受け取り接続が閉じられた状態 |
| UNKNOWN | 状態不明のソケット |

### curl

### wget

### tail

### head

### less

### find

### ssh

### kill

## Memory Management

## Interprocess Communication

## I/O Management

## POSIX Basics

### stdin

### stdout

### stdrr

### pipes

## Basic Networking Concepts
