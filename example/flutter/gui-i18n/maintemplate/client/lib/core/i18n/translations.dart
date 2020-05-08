import 'package:flutter/material.dart';
import 'package:intl/intl.dart';

class Translations {
  Locale locale;

  String testMessage() {
    return Intl.message(
      'This is a i18n test message. If you can read it in different languages after pressing one of the buttons below everything should be setup correctly.',
      name: 'testMessage',
      desc: 'test example for i18n tool',
      locale: locale.toString(),
    );
  }

  String testMessageTwo() {
    return Intl.message(
      'Hey here is the second test String... : - ) Let\'s see if this works',
      name: 'testMessageTwo',
      desc: 'test example for i18n tool',
      locale: locale.toString(),
    );
  }

  String testMessageThree() {
    return Intl.message(
      'Third one! Hey google translate, what\'s up?',
      name: 'testMessageThree',
      desc: 'test example for i18n tool',
      locale: locale.toString(),
    );
  }
}
