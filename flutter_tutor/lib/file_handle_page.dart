import 'dart:io';
import 'dart:async';

import 'package:file_selector/file_selector.dart';
import 'package:flutter/foundation.dart';
import 'package:flutter/material.dart';
import 'package:path_provider/path_provider.dart';
import 'package:file_picker/file_picker.dart';

class MyFileStorage {
  Future<String> get _localPath async {
    final directory = await getApplicationDocumentsDirectory();

    if (kDebugMode) {
      print(directory.path);
    }

    return directory.path;
  }

  Future<File> get _localFile async {
    final path = await _localPath;

    if (kDebugMode) {
      print('$path/swgkeys.txt');
    }

    return File('$path/swgkeys.txt');
  }

  Future<int> readCounter() async {
    try {
      final file = await _localFile;
      final contents = await file.readAsString();

      return int.parse(contents);
    } catch(e) {
      return 0;
    }
  }

  Future<File> writeCounter(int counter) async {
    final file = await _localFile;

    return file.writeAsString('$counter');
  }
}

class FileHandlePage extends StatefulWidget {
  final MyFileStorage storage;

  const FileHandlePage({
    Key? key,
    required this.storage,
  }) : super(key: key);

  @override
  State<FileHandlePage> createState() => _FileHandlePageState();
}

class _FileHandlePageState extends State<FileHandlePage> {
  int _counter = 0;

  @override
  void initState() {
    super.initState();
    widget.storage.readCounter()
        .then((value) => setState(() {
          _counter = value;
        }));
  }

  Future<File> _incrementCounter() {
    setState(() {
      _counter++;
    });

    return widget.storage.writeCounter(_counter);
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Reading and writing Files'),
      ),
      body: Center(
        child: Text(
          'Button tapped $_counter time${_counter == 1 ? '' : 's'}.',
        ),
      ),
      floatingActionButton: FloatingActionButton(
        onPressed: _incrementCounter,
        tooltip: 'Increment',
        child: const Icon(Icons.add),
      ),
    );
  }
}

class LocalFileHandlePage extends StatefulWidget {
  String filePath = '';
  final xType = const XTypeGroup(label: '视频', extensions: ['mp4', 'rm', 'mpg']);

  LocalFileHandlePage({
    Key? key,
    required this.filePath,
  }) : super(key: key);

  @override
  State<LocalFileHandlePage> createState() => _LocalFileHandlePageState();
}

class _LocalFileHandlePageState extends State<LocalFileHandlePage> {
  String path = '';

  @override
  void initState() {
    super.initState();
    setState(() {
      path = widget.filePath;
    });
  }

  @override
  Widget build(BuildContext context) {
    return Container();
  }
}
