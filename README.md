
This is just a simple command line tool to format a list of URLS into a tree for review.

#### urls.txt:
```
https://www.test.com/content/assets/css/colorbox.css
https://www.test.com/content/assets/css/custom.css
https://www.test.com/data/dbllhv900001412-1028463?index=7
https://www.test.com/data/dbllhv900001865-0039053?index=8
https://www.test.com:8443/
https://www.test.com:8443/content/assets/css/code.css
https://www.test.com:8443/content/assets/css/rst.css
https://www.test.com:8443/content/assets/css/theme.css
https://www.test.com:8443/?page=1&sort=dec
https://www.test.com:8443/?page=1&sort=galaxy
https://www.test.com:8443/?page=1&sort=data_name
https://www.test.com:8443/?page=1&sort=ra
https://www.test.com:8443/?page=1&sort=lmr_id
https://www.test.com:8443/?page=7672&sort=d26
https://www.test.com:8443/?page=7672&sort=dec
https://www.test.com:8443/?page=7672&sort=galaxy
https://www.test.com:8443/?page=7672&sort=data_name
https://www.test.com:8443/?page=7672&sort=lmr_id
https://www.test.com:8443/?page=7673&sort=d26
https://www.test.com:8443/?page=7673&sort=dec
https://www.test.com:8443/?page=7673&sort=galaxy
https://www.test.com:8443/?page=7673&sort=data_name
https://www.test.com:8443/?page=7673&sort=ra
https://www.test.com:8443/?page=7673&sort=lmr_id
https://www.test.com:8443/data/dbllhv900001412-1028463?index=7
https://www.test.com:8443/data/dbllhv900001865-0039053?index=8
https://www.test.com:8443/data/dbllhv900041319+1433044?index=22
https://www.test.com:8443/data/DR8-2585p145-4805?index=9
https://www.test.com:8443/data/ESO550-023?index=12
https://www.test.com:8443/data/NGC3031
https://www.test.com:8443/data/fnZZ172431?index=383599
https://www.test.com:8443/data/fnZZ66749?index=43
https://www.test.com:8443/data/fnZZ26151?index=383601
https://blah.gooble.com/
https://blah.gooble.com/static/assets/css/code.css
https://blah.gooble.com/static/assets/css/colorbox.css
https://blah.gooble.com/static/assets/css/custom.css
https://blah.gooble.com/static/assets/css/rst.css
https://blah.gooble.com/static/assets/css/theme.css
https://blah.gooble.com/?page=1&sort=dec
https://blah.gooble.com/?page=1&sort=galaxy
https://blah.gooble.com/?page=1&sort=group_name
https://blah.gooble.com/?page=1&sort=ra
https://blah.gooble.com/?page=1&sort=lmr_id
https://blah.gooble.com/?page=7672&sort=d26
https://blah.gooble.com/?page=7672&sort=dec
https://blah.gooble.com/?page=7672&sort=bha
https://blah.gooble.com/?page=7672&sort=group_name
https://blah.gooble.com/?page=7672&sort=test
https://blah.gooble.com/?page=7673&sort=d26
https://blah.gooble.com/?page=7673&sort=dec
https://blah.gooble.com/?page=7673&sort=galaxy
https://blah.gooble.com/?page=7673&sort=group_name
https://blah.gooble.com/?page=7673&sort=ra
https://blah.gooble.com/?page=7673&sort=lmr_id
https://blah.gooble.com/group/2MASXJ00001412-1028463?index=7
https://blah.gooble.com/group/2MASXJ00001865-0039053?index=8
https://blah.gooble.com/group/2MASXJ00041319+1433044?index=22
https://blah.gooble.com/group/DR8-2585p145-4805?index=9
https://blah.gooble.com/group/ESO550-023?index=12
https://blah.gooble.com/group/NGC3031
https://blah.gooble.com/group/fnZZ172431?index=383599
https://blah.gooble.com/group/fnZZ66749?index=43
https://blah.gooble.com/group/fnZZ26151?index=383601
```

#### output
```
cat ./urls.txt | urltree 
https://www.test.com:443
├── content
│   └── assets
│       └── css
│           ├── colorbox.css
│           └── custom.css
└── data
    ├── dbllhv900001412-1028463
    └── dbllhv900001865-0039053

https://www.test.com:8443
├── content
│   └── assets
│       └── css
│           ├── code.css
│           ├── rst.css
│           └── theme.css
└── data
    ├── DR8-2585p145-4805
    ├── ESO550-023
    ├── NGC3031
    ├── dbllhv900001412-1028463
    ├── dbllhv900001865-0039053
    ├── dbllhv900041319+1433044
    ├── fnZZ172431
    ├── fnZZ26151
    └── fnZZ66749

https://blah.gooble.com:443
├── group
│   ├── 2MASXJ00001412-1028463
│   ├── 2MASXJ00001865-0039053
│   ├── 2MASXJ00041319+1433044
│   ├── DR8-2585p145-4805
│   ├── ESO550-023
│   ├── NGC3031
│   ├── fnZZ172431
│   ├── fnZZ26151
│   └── fnZZ66749
└── static
    └── assets
        └── css
            ├── code.css
            ├── colorbox.css
            ├── custom.css
            ├── rst.css
            └── theme.css
```