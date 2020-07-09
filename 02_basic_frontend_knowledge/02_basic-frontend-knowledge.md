# HTML

HyperText Markup Language.
Webページを作成するために開発された言語。
ハイパーテキストに目印を付ける言語。ハイパーテキストの説明はintroでやったので省略。

文書の各部分が、どのような役割を持っているのかを示す。
文書内の各部分に目印をつけて、その部分がどんな要素なのかを明確にすることで、コンピュータがその文書の構造を理解出来るようになる。

## ブロックレベル要素

ブロックレベル要素は、見出し・段落・表・など、文書を構成する基本となる要素で、一つのブロックとして認識される。
ブラウザでの表示も一つのかたまりとして扱われることが多く、一般的なブラウザでは前後に改行が入る。以下実例。

`<address>、<blockquote>、<center>、<div>、<dl>、<fieldset>、<form>、<h1>-<h6>、<hr>、<noframes>、<noscript>、<ol>、<p>、<pre>、<table>、<ul>`

## インライン要素

インライン要素は、主にブロックレベル要素の内容として用いられている要素で、文章の一部として扱わる。
例えば、`<p>`要素の中の`<strong>`要素のように、段落のなかの一部を強調するような使われ方をする要素。
一般的なブラウザでは前後に改行は入らず、文章の一部として表示される。

`<a>、<abbr>、<acronym>、<b>、<basefont>、<bdo>、<big>、<br>、<cite>、<code>、<dfn>、<em>、<font>、<i>、<img>、<input>、<kbd>、<label>、<q>、<s>、<samp>、<select>、<small>、<span>、<strike>、<strong>、<sub>、<sup>、<textarea>、<tt>、<u>、<var>`

ブロックレベル要素の中にインライン要素を入れることはできません😷

## 汎用属性

|  属性名  |  内容  |
| ---- | ---- |
|  style  |  要素に直接スタイルシートを適用する際に使用する。一部だけスタイルを当てたい時に便利。  |
|  class  |  要素にクラス名を付ける際に使用。クラス名は1つの文書内で重複して同じ名前を指定することが出来る。
              文書内の複数箇所に同じスタイルを適用する場合に便利な属性。
              スタイルの役割で命名するといい。  |
|  id  |  要素に固有の名前を付ける際に使用すえう。id名はファイル内で一意。
          そのため、id名を付与した箇所を一意に特定して、スタイルシートやスクリプトを適用することが出来る。  |
|  title  |  要素の補足情報を指定する際に使用。ツールチップとして表示されることが多い。  |
|  lang  |  要素内容の言語を指定する際に使用する。
            例えば、日本語の文章中に英文が混在する場合には、音声ブラウザなどへの影響を考慮してlang属性を指定する。  |
|  dir  |  要素内容の表示方向や読み上げ方向を指定する際に使用する。
            英語や横書きの日本語のように、左から右へ表示したい場合はltrを、右から左へ表示したい場合はrtlを指定する。  |

## サイズ指定

- ピクセル ： 言わずとしれたピクセル表記。スクリーンの1pxの長さを1とした単位。実際にモニタに表示されるサイズは、解像度に準拠。
- % : 水平・垂直方向の表示可能な領域に対する割合で指定する。
- 長さ比率 : アスタリスクを使用して、長さ比率を指定する。
  - 長さを複数に分割する場合などに用いる。例えば「`*,3*`」と記述した場合、長さを1:3の比率で分割する。
  - ピクセル数や%による長さ指定を併用している場合は、ピクセル数や％で指定した分が先に確保され、残りの長さが*で指定した比率で分割される。

## CSS

略

## JavaScript

略