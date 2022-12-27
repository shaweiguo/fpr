import 'dart:io';
import 'package:flutter/foundation.dart';
import 'package:flutter/material.dart';
import 'package:video_player/video_player.dart';
import 'package:file_picker/file_picker.dart';
// import 'package:open_file/open_file.dart';

class VideoPlayerPage extends StatefulWidget {
  const VideoPlayerPage({Key? key}) : super(key: key);

  @override
  State<VideoPlayerPage> createState() => _VideoPlayerPageState();
}

class _VideoPlayerPageState extends State<VideoPlayerPage> {
  late VideoPlayerController _controller;
  String _filePath  = 'assets/videos/20.mp4';

  // void _loadVideoPlayer() {
  //   _controller = VideoPlayerController.asset('assets/videos/01.mp4');
  //   _controller.addListener(() {
  //     setState(() {
  //       // TODO: Add events handle
  //     });
  //   });
  //   _controller.initialize().then((value) {
  //     setState(() {
  //       //
  //     });
  //   });
  // }

  @override
  void initState() {
    super.initState();
    _loadVideo();
    // _loadVideoPlayer();
  }

  void _pickFile() async {
    final result = await FilePicker.platform.pickFiles(allowMultiple:  false);

    if (result == null) return;

    if (kDebugMode) {
      print(result.files.first.name);
      print(result.files.first.size);
      print(result.files.first.path);
    }

    setState(() {
      _filePath = result.files.first.path!;
      if (_controller.value.isPlaying) {
        _controller.seekTo(Duration.zero);
      }
      _loadVideo();
    });
  }

  void _loadVideo() {
    _controller = VideoPlayerController.file(File(_filePath))
      ..setVolume(1.0)
      ..initialize()
      ..play();
  }
  
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Play Video'),
        backgroundColor: Colors.redAccent,
      ),
      body: Column(
        children: [
          IconButton(
            onPressed: () {
              _pickFile();
            },
            icon: const Icon(Icons.file_open),
          ),
          // _controller.value.isInitialized ?
            AspectRatio(
              aspectRatio: _controller.value.aspectRatio,
              child: VideoPlayer(_controller),
            ), // :
            // Container(),
          Text("Total Duration: ${_controller.value.duration.toString()}"),
          VideoProgressIndicator(
            _controller,
            allowScrubbing: true,
            colors: const VideoProgressColors(
              backgroundColor: Colors.redAccent,
              playedColor: Colors.green,
              bufferedColor: Colors.purple,
            ),
          ),
          Row(
            children: [
              IconButton(
                onPressed: () {
                  if (_controller.value.isPlaying) {
                    _controller.pause();
                  } else {
                    _controller.play();
                  }
                  setState(() {
                    // TODO:
                  });
                },
                icon: Icon(
                  _controller.value.isPlaying ? Icons.pause : Icons.play_arrow
                ),
              ),
              IconButton(
                onPressed: () {
                  _controller.seekTo(const Duration(seconds: 0));
                },
                icon: const Icon(Icons.stop),
              ),
            ],
          )
        ],
      ),
    );
  }

  @override
  void dispose() {
    super.dispose();
    _controller.dispose();
  }
}
