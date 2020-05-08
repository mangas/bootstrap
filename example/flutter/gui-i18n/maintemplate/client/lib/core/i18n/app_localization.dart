import 'package:flutter/cupertino.dart';
import 'package:flutter/foundation.dart';
import 'package:maintemplate/core/i18n/generated/messages_all.dart';
import 'package:maintemplate/core/i18n/translations.dart';
import 'package:flutter/widgets.dart';

class AppLocalizations extends Translations {
  final Locale locale;
  static Map<String, String> supportedLanguages = {
    'en': 'English',
    'fr': 'French',
    'de': 'German',
    'es': 'Spanish',
    'it': 'Italian',
    'ur': 'Urdu',
  };

  static List<Locale> getLocales() {
    List<Locale> _locales = [];

    supportedLanguages.forEach((key, value) {
      _locales.add(Locale(key));
    });

    return _locales;
  }

  AppLocalizations(this.locale);

  static Future<AppLocalizations> load(Locale locale) {
    return initializeMessages(locale.toString()).then((Object _) {
      return new AppLocalizations(locale);
    });
  }

  static AppLocalizations of(BuildContext context) {
    return Localizations.of<AppLocalizations>(context, AppLocalizations);
  }
}

class AppLocalizationsDelegate extends LocalizationsDelegate<AppLocalizations> {
  final Locale overriddenLocale;

  AppLocalizationsDelegate(this.overriddenLocale);

  @override
  bool shouldReload(AppLocalizationsDelegate old) => true;

  @override
  bool isSupported(Locale locale) {
    return AppLocalizations.supportedLanguages.keys
        .contains(locale.languageCode);
  }

  @override
  Future<AppLocalizations> load(Locale locale) {
    //if (this.overriddenLocale == Locale('en')) {
    //  print("return system");
    //  return AppLocalizations.load(locale);
    //}
    //print("return overriden");
    return AppLocalizations.load(locale);
  }
}

class FallbackCupertinoLocalisationsDelegate
    extends LocalizationsDelegate<CupertinoLocalizations> {
  const FallbackCupertinoLocalisationsDelegate();

  @override
  bool isSupported(Locale locale) =>
      AppLocalizations.supportedLanguages.keys.contains(locale.languageCode);

  @override
  Future<CupertinoLocalizations> load(Locale locale) =>
      SynchronousFuture<_DefaultCupertinoLocalizations>(
          _DefaultCupertinoLocalizations(locale));

  @override
  bool shouldReload(FallbackCupertinoLocalisationsDelegate old) => false;
}

class _DefaultCupertinoLocalizations extends DefaultCupertinoLocalizations {
  final Locale locale;

  _DefaultCupertinoLocalizations(this.locale);
}
