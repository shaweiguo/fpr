import 'dart:io';
import 'package:file_selector/file_selector.dart';
import 'package:flutter/material.dart';
import 'package:file_selector_platform_interface/file_selector_platform_interface.dart';
import 'package:video_player/video_player.dart';

class OpenVideoPage extends StatelessWidget {
  const OpenVideoPage({Key? key}) : super(key: key);

  Future<void> _openVideoFile(BuildContext context) async {
    const XTypeGroup typeGroup = XTypeGroup(
      label: 'videos',
      extensions: <String>['mp4', 'rm', 'avi'],
    );
    final XFile? file = await FileSelectorPlatform.instance
      .openFile(acceptedTypeGroups: <XTypeGroup>[typeGroup]);

    if (file == null) {
      print("file is null");
      return;
    }

    final String fileName = file.name;
    final String filePath = file.path;

    await showDialog(
      context: context,
      builder: (BuildContext context) =>  VideoPlayerPage(fileName, filePath),
    );
  }

  @override
  Widget build(BuildContext context) {
    return Container();
  }
}

class VideoPlayerPage extends StatefulWidget {
  final String fileName;
  final String filePath;
  const VideoPlayerPage(this.fileName, this.filePath, {Key? key}) : super(key: key);
  
  @override
  State<VideoPlayerPage> createState() => _VideoPlayerPageState();
}

class _VideoPlayerPageState extends State<VideoPlayerPage> {
  late VideoPlayerController _controller;
  
  @override
  void initState() {
    super.initState();
    _controller = VideoPlayerController.file(File(widget.filePath))
      ..initialize().then((_) => setState(() {}));
  }

  @override
  Widget build(BuildContext context) {
    return Column(
      children: [
        _controller.value.isInitialized
          ? AspectRatio(
              aspectRatio: _controller.value.aspectRatio,
              child: VideoPlayer(_controller),
            )
          : Container(),
        IconButton(
          onPressed: () {},
          icon: _controller.value.isPlaying
            ? const Icon(Icons.pause)
            : const Icon(Icons.play_arrow)
        ),
      ],
    );
  }

  @override
  void dispose() {
    super.dispose();
    _controller.dispose();
  }
}
