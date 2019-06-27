package wasm

import hex "github.com/tmthrgd/go-hex"

var Bytecode_storage_test = hex.MustDecodeString("0061736D0100000001220760037F7F7F0060027F7F0060000060017F0060017F017F60027F7F017F6000017E02380303656E760D7365745F73746F726167653332000003656E760D6765745F73746F726167653332000003656E76066D656D6F727902010202030F0E00010102030405000002060203040608017F01418080040B07BB010E095F5F6D656D637079380002085F5F627A65726F380003075F5F627365743800040B5F5F696E69745F686561700005065F5F667265650006085F5F6D616C6C6F630007095F5F7265616C6C6F6300080B5F5F62653332746F6C654E00090B5F5F6C654E746F62653332000A12736F6C3A3A5F5F636F6E7374727563746F72000B10736F6C3A3A676574466F6F506C757332000C0B736F6C3A3A696E63466F6F000D0B636F6E7374727563746F72000E0866756E6374696F6E000F0AF8080E2600034020002001290300370300200041086A2100200141086A21012002417F6A22020D000B0B1C00034020004200370300200041086A21002001417F6A22010D000B0B1C0003402000427F370300200041086A21002001417F6A22010D000B0B2B00410041003602808004410041003602848004410041003A008C800441003F0041F0FF7B6A36028880040BB90101037F02402000450D00410021012000417C6A41003A0000024002400240200041706A22022802002203450D0020032D000C450D01200321010B200041746A2802002203450D020C010B20022003280200220136020002402001450D00200120023602040B200041786A2202200328020820022802006A41106A360200200041746A2802002203450D010B20032D000C0D0020012003360204200320013602002003200041786A28020020032802086A41106A3602080F0B0BA90101047F41808004210102400340024020012D000C0D002001280208220220004F0D020B200128020022010D000B41002101410028020821020B02402002200041076A41787122036B22024118490D00200120036A41106A22002001280200220436020002402004450D00200420003602040B2000200241706A360208200041003A000C2000200136020420012000360200200141086A20033602000B200141013A000C200141106A0BF50101057F024002400240200041706A22022802002203450D0020032D000C0D00200041786A220428020020032802086A41106A220520014F0D010B41002103410020014103766B21022001100721010340200120036A200020036A290300370300200341086A2103200241016A22020D000B20001006200121000C010B20042005360200200328020022032002360204200220033602002005200141076A41787122066B22054118490D00200020066A2201200336020002402003450D00200341046A20013602000B2001200541706A360208200141003A000C20012002360204200220013602002004200636020020000F0B20000B2D002000411F6A21000340200120002D00003A0000200141016A21012000417F6A21002002417F6A22020D000B0B2D002001411F6A21010340200120002D00003A00002001417F6A2101200041016A21002002417F6A22020D000B0B2701017F230041106B22002400200042E6003703084100200041086A41081000200041106A24000B2D02017F017E230041106B220024004100200041086A4108100120002903082101200041106A2400200142027C0B3701017F230041106B220024004100200041086A410810012000200029030842017C3703084100200041086A41081000200041106A24000B12001005024020002802000D00100B0F0B000B8E0103037F017E017F23002201210202402000280200220341034D0D002003417C6A21030240200041046A28020041DE9AB88F79470D0020030D01100D41041007220041003602002002240020000F0B20030D00100C21044124100722004120360200200041046A220541041003200141706A2203240020032004370300200320054108100A2002240020000F0B000B")
var Abi_storage_test = []byte(`[{"name":"","type":"constructor","inputs":[],"outputs":[],"constant":false,"payable":false,"stateMutability":"nonpayable"},{"name":"getFooPlus2","type":"function","inputs":[],"outputs":[{"name":"","type":"uint64"}],"constant":true,"payable":false,"stateMutability":"view"},{"name":"incFoo","type":"function","inputs":[],"outputs":[],"constant":false,"payable":false,"stateMutability":"nonpayable"}]`)