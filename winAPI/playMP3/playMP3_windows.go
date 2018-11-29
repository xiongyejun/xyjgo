package playMP3

import (
	"errors"

	"github.com/xiongyejun/xyjgo/winAPI/winmm"
)

// 释放dll
func Free() {
	winmm.Free()
}

// 播放MP3
// MP3File	MP3文件路径
// bWait	是否等待播放结束
func Play(MP3File string, bWait bool) (err error) {
	Stop()

	strCMD := "open \"" + MP3File + "\" alias mymusic"
	ret := winmm.MciSendStringA(strCMD, "", 0, 0)
	if ret != 0 {
		return errors.New("open " + winmm.GetErrorString(ret))
	}

	if bWait {
		ret = winmm.MciSendStringA("play mymusic wait", "", 0, 0)
	} else {
		ret = winmm.MciSendStringA("play mymusic", "", 0, 0)
	}

	if ret != 0 {
		return errors.New(winmm.GetErrorString(ret))
	}

	return nil
}

// 停止播放
func Stop() {
	winmm.MciSendStringA("stop mymusic", "", 0, 0)
	winmm.MciSendStringA("close mymusic", "", 0, 0)
}

// 暂停播放
func Pause() (err error) {
	ret := winmm.MciSendStringA("pause mymusic", "", 0, 0)
	if ret != 0 {
		return errors.New(winmm.GetErrorString(ret))
	}
	return nil
}

// 继续播放
func Continue() (err error) {
	ret := winmm.MciSendStringA("play mymusic", "", 0, 0)
	if ret != 0 {
		return errors.New(winmm.GetErrorString(ret))
	}
	return nil
}

//func Pause() bool {
//	return 0 == winmm.MciSendStringA(syscall.StringToUTF16Ptr("pause music"), syscall.StringToUTF16Ptr(""), 0, 0)
//}

//func Continue() bool {
//	return 0 == winmm.MciSendStringA(syscall.StringToUTF16Ptr("play music"), syscall.StringToUTF16Ptr(""), 0, 0)
//}

//    Function PlayMidiFile(ByVal MusicFile As String) As Boolean
//        mciSendStringA("stop music", "", 0, 0)
//        mciSendStringA("close music", "", 0, 0)
//        mciSendStringA("open " & MusicFile & " alias music", "", 0, 0)
//        PlayMidiFile = mciSendStringA("play music", "", 0, 0) = 0
//    End Function

//    Function StopMidi() As Boolean
//        StopMidi = mciSendStringA("stop music", "", 0, 0) = 0
//        mciSendStringA("close music", "", 0, 0)
//    End Function

//    Private Function PauseMidi() As Boolean
//        Return mciSendStringA("pause music", "", 0, 0) = 0
//    End Function

//    Private Function ContinueMidi() As Boolean
//        Return mciSendStringA("play music", "", 0, 0) = 0
//    End Function
