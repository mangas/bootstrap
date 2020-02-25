

import 'package:flutter/cupertino.dart';
import 'package:intl/intl.dart';
import 'package:flutter/widgets.dart';

import 'generated/mock_text_messages_all.dart';

class AppLocalizations{
  final Locale locale;


  AppLocalizations(this.locale);

  static const LocalizationsDelegate<AppLocalizations> delegate = _AppLocalizationsDelegate();

  static Future<AppLocalizations> load(Locale locale) {
    return initializeMessages(locale.toString())
      .then((Object _) {
        return new AppLocalizations(locale);
      });
  }

  static AppLocalizations of(BuildContext context) {
    return Localizations.of<AppLocalizations>(context, AppLocalizations);
  }

  String mockText() {
    return Intl.message(
      'Im the text to get translated, change language in your phone settings, and try me',
      name: 'mockText',
      desc: 'mock text for the application',
      locale: locale.toString(),
    );
  }

}


class _AppLocalizationsDelegate extends LocalizationsDelegate<AppLocalizations> {

  const _AppLocalizationsDelegate();

  

  @override
  bool shouldReload(_AppLocalizationsDelegate old) => false;

  @override
  bool isSupported(Locale locale) {
    return ['en', 'es', 'fr', 'ur'].contains(locale.languageCode);
  }

  @override
  Future<AppLocalizations> load(Locale locale) => AppLocalizations.load(locale);

}