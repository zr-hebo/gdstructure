# gdstructure
gdstructure是一个golang实现的常用数据结构

# Install

Use go get to install this package:<br>
$ go get github.com/zr-hebo/gdstructure

# Usage

  import (
    bfw "github.com/zr-hebo/bufferedwriter"
  )

	// 创建带缓存的Writer
  // file是打开的文件
  // bufferSize是缓存的大小
  writer, err := bfw.NewBufferedWriter(file, bufferSize)
	if err != nil {
		return
	}
  
  // 将byte数组写入writer	
  if err = writer.Write(rowData); err != nil {
		return
	}

	// 等待缓存中的数据全部写入文件
	if err = writer.WaitClean(); err != nil {
		return
	}

