# gowins
golang for windows


## NonClientMetrics

[lxn/walk](https://github.com/lxn/walk)のデフォルトフォントがシステムフォントになっていないみたいなので、SystemParametersInfoから取得するための処理。

システム設定のDPIとか変更した際には再度`Get()`で取得しなおす必要があります。

> たしか変更された際にはWndProcとかで通知されるはず。


## SingleInstanceApplication

多重起動チェック用で、内部でCreateMutexを使用。


## WindowExStyle

コントロールがウィンドウ全体を使う場合だとエッジで見栄えが悪くなるので、
ClientEdgeはWS_EX_CLIENTEDGEの有効無効を設定できるようにしています。
