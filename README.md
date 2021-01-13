# Dev

```
git clone https://github.com/cclin81922/go-flutter.git
cd go-flutter

cd fe
flutter build web
cd ..

cd fe_fetch_data_from_the_internet
flutter build web
cd..

cd fe_web_dashboard
flutter build web
cd ..

go build
./go-flutter
```

# Install

```
go get github.com/cclin81922/go-flutter
~/go/bin/go-flutter
```

# Usage

```
go-flutter
go-flutter --port 8080
go-flutter --fe fe/build/web
go-flutter --fe fe_web_dashboard/build/web
go-flutter --fe fe_fetch_data_from_the_internet/build/web
```
