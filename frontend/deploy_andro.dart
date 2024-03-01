import 'dart:async';
import 'dart:io';

const url = 'localhost:8080/api/upload';
const apkPath = 'build/app/outputs/flutter-apk/app-release.apk';

const flutterPathWindows =
    'C:\\Users\\user\\Documents\\flutter\\bin\\flutter.bat';

const flutterPathMac = '/Users/gspe/developments/flutter/bin/flutter';

String functionName = '';

Future<void> main() async {
  final flutterPath = Platform.isWindows ? flutterPathWindows : flutterPathMac;
  final Timer spinner = startLoadingSpinner();

  try {} catch (err) {
    print("Error occured: $err");
  }

  /// dispose Spinner timer
  spinner.cancel();
}

Timer startLoadingSpinner() {
  // const spinnerChars = ['|', '/', '-', '\\'];
  const spinnerSequence = ['⣷', '⣯', '⣟', '⡿', '⢿', '⣻', '⣽', '⣾'];

  int index = 0;

  return Timer.periodic(const Duration(milliseconds: 100), (timer) {
    stdout.write(
        '\r$functionName${'...'.padRight(55, ' ')} ${spinnerSequence[index]}');
    index = (index + 1) % spinnerSequence.length;
  });
}

Future<String> getCurrentVersion() async {
  final pubspecFile = File('pubspec.yaml');
  final lines = await pubspecFile.readAsLines();
  for (final line in lines) {
    if (line.startsWith('version:')) {
      return line.split(':')[1].trim();
    }
  }
  throw 'Version not found in pubspec.yaml';
}