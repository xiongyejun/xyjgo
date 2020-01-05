/*
Package mp4 implements encoding/decoding of MP4 media.
*/
package mp4

/*
https://www.cnblogs.com/ranson7zop/p/7889272.html
http://www.52rd.com/Blog/wqyuwss/559/4/

rack  表示一些sample的集合，对于媒体数据来说，track表示一个视频或音频序列。
hint track  这个特殊的track并不包含媒体数据，而是包含了一些将其他数据track打包成流媒体的指示信息。
sample  对于非hint track来说，video sample即为一帧视频，或一组连续视频帧，audio sample即为一段连续的压缩音频，它们统称sample。对于hint track，sample定义一个或多个流媒体包的格式。
sample table  指明sampe时序和物理布局的表。
chunk 一个track的几个sample组成的单元。
*/
