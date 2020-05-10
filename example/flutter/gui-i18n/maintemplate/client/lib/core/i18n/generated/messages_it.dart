// DO NOT EDIT. This is code generated via package:intl/generate_localized.dart
// This is a library that provides messages for a it locale. All the
// messages from the main program should be duplicated here with the same
// function name.

// Ignore issues from commonly used lints in this file.
// ignore_for_file:unnecessary_brace_in_string_interps, unnecessary_new
// ignore_for_file:prefer_single_quotes,comment_references, directives_ordering
// ignore_for_file:annotate_overrides,prefer_generic_function_type_aliases
// ignore_for_file:unused_import, file_names

import 'package:intl/intl.dart';
import 'package:intl/message_lookup_by_library.dart';

final messages = new MessageLookup();

typedef String MessageIfAbsent(String messageStr, List<dynamic> args);

class MessageLookup extends MessageLookupByLibrary {
  String get localeName => 'it';

  final messages = _notInlinedMessages(_notInlinedMessages);
  static _notInlinedMessages(_) => <String, Function> {
    "testMessage" : MessageLookupByLibrary.simpleMessage("Questo è un messaggio di prova i18n. Se riesci a leggerlo in diverse lingue dopo aver premuto uno dei pulsanti in basso, tutto dovrebbe essere impostato correttamente."),
    "testMessageThree" : MessageLookupByLibrary.simpleMessage("Terza! Ehi, google traduttore, come va?"),
    "testMessageTwo" : MessageLookupByLibrary.simpleMessage("Ehi, ecco la seconda stringa di test ...: -) Vediamo se funziona")
  };
}
