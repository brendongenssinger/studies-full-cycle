#!/bin/bash

# Lista de volumes a serem removidos
volumes=(
  "2b0d3094d332e8f00123f2f8935f792db73955cbd62b31348831e7d18a9b1306"
  "4d8f85eff3b65262370496ca696b3c2e69d1e9c45c2684c7f43e2e48d9abbad0"
  "5bd8a231fff0aaa836fe032e854330d681511ef8e8d75cba6277ce66e786b230"
  "6a08888b3b51b31014ddf44504b09d086cbc969dee8618d8c24d18458fa3b651"
  "7ec5a05dcb5b7a53f56915dc2ed9d6ecca2ca62681c79ba16c1cc1371ce825e8"
  "8dde2c20cd06181bc2154a639bba9540793dd0d6dca18d96bae19f75e1df5e8c"
  "9a0e7c324bc97f968f18da1d2037fb5e00cffef7c59fbfcd8b99a3c079590529"
  "9c970d3696c60b9f45bdf4d26d5a7c08d280e4ccf9941bc4df82754a95cdc2b7"
  "46a3a04a41903b4aa01956c4532dd135a98f42d1aca753074c0080f7863d393f"
  "64bc0b0a25336e3be8adf7fb68f1e28667f987cc0b47991c110e3b5dc227f2ba"
  "591e39a547da01a090637c7dec2e2dcdb2a8c46d81abc08094bc89bbc3701494"
  "835bbabc61be3fca20122c0d19482eebdb81a3f040d78f74b80bb469bd51307b"
  "1381e08d97ef1c70385280db0a2e23533c76ee35320252a37804205d571a9253"
  "10007e4b32fabe3717054316b7cff4ac8e49c8a6774f4122e3368d66ee83d2e7"
  "65993cf779e3ce8934f82cf82ed323e0eeb41c0dce26ca83227e2a452726c4b7"
  "875367dff17952052dc95e97127c5f89a7b7a060604e973af419e63c95719118"
  "42835152c6a98c9aee8e35cebe372b29f0c8566a8518fc4c69c8743153683b1b"
  "468145942b0c66ab840d0552e56ef70838a41a20be96f0bf410370ea07def6d2"
  "711778025c0175055f00c6b9d4dbf44d36f430a4209e41acce67f7fbb0550df7"
  "2496857540a2943a195090279cc3400e597c5532071c710da282e1cfaf469695"
  "a6af96fc0604545b172dd476723b00899c9c773acd0cace47c2b46efbaf9a137"
  "a9ba7af6e7677b8fa7accd0ffcece75adb7c5f145ace563150340d769cf583d9"
  "afba2abda71669cb13810211c098ecaf425203eddaab44e9a962cd1c3e334dd0"
  "b75c9a104d0ed176d4a41ec60c2e02d0e341f05f56cf63c36e4bb9292746f4fa"
  "bdd1b83dbd543f4990eeca8330887c0d9da070b8c14f15d48f72ebc6fcc7b009"
  "bf7ac7fcd4f7002faf0651ff381a2181482cc0cd45580dfb9fc11f838f565b5d"
  "c64349b87e456c9dda611b4a5585acbef5c087c673e1e5774de14ff77bacbd25"
  "d03c2d2c175fefb421d6cefa1426690e13f657c345da72ca16837b0aca1ce017"
  "d206ddfff9818a0d55cd03482ad8e2221244357ec106b57379ebb3f22b4f9026"
  "d474ae0828794807b811bc047202db16b0b9af7acccfc6f416ae45ac674c3089"
  "e9e84f91008d6c7587c5cd814a6ec2cdb15e8b0906517c0aea9325b9063a9209"
  "e3993a7dc9fed2623f1af7c4354b8e3178ab8fe784e44ea260dc8986c6084f6c"
  "ec1f37ed3583d9ad226dec638df6dfdd07a29f8f961c45699d0074c2e5e031e2"
  "meuvolume"
)

# Loop para remover cada volume
for volume in "${volumes[@]}"; do
  docker volume rm "$volume"
done