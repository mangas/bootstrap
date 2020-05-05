# i18n_example

This is the i18n tool Flutter example.

It won't compile until you add the translations for the first time.

# 1. Add new strings or change old ones
To add new translation in the Flutter project, do it in lib/core/i18n/translations.dart

# 2. Transform strings from dart code into .arb file and translate them
This is done with the Makefile

`make lang-gen-flu`
or manually:

``` 
mkdir -p lib/core/i18n/generated/
mkdir -p i18n/
flutter pub run intl_translation:extract_to_arb --output-dir=i18n/ lib/core/i18n/translations.dart
i18n flutter --dir i18n/ --template i18n/intl_messages.arb --prefix lang --languages en,fr,es,de,it,ur -f
i18n flutter --dir i18n/
```

# 3. Add translated strings back again to code!

`make lang-gen-flu-dart`
or manually:

```
flutter pub run intl_translation:generate_from_arb --output-dir=lib/core/i18n/generated/ lib/core/i18n/translations.dart i18n/*.arb
```

# 4. Start Flutter project and test
e.g.

flutter pub get
flutter run -d chrome
