// Package spec provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package spec

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	externalRef0 "github.com/trustbloc/vcs/pkg/restapi/v1/common"
)

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/+x963IbN9bgq6D4bVXsWpKyc5n5ov3zKZKTKGPH+iTZrqnYxYK6QRJWs9EB0KI5KW/t",
	"a+zr7ZNs4QDoBrqBvlCSY0/0ZyYWG7eDcw7O/fwxSdimYDnJpZgc/jERyZpsMPznUZIQIS7ZNcnPiShY",
	"Loj6c0pEwmkhKcsnh5MXLCUZWjKO9OcIvkd2wHwynRScFYRLSmBWDJ8tpPqsPd3lmiD9BYIvEBWiJCm6",
	"2iGpfirlmnH6L6w+R4LwG8LVEnJXkMnhREhO89Xk43TifbhIicQ0E+3lzp/996vT82cnaLsmOQoOQgXm",
	"eEMk4YgKVAqSIskQJ7+XREjYHs4TgtgSYZQQLjHN0TEnKcklxRlSO0NYoJQsaU5SRHN0QRLY/nfzp/On",
	"c3Qq0YtXF5fo15eX6IroFZhcE76lgsDPVCCcI8w53ql12NV7kkgxjUz7d/XNb+c/Hn//zfd/e6egQyXZ",
	"wOH/ByfLyeHkPw7qOz8wF67+tGH54siFwYmB28cKwrAJ9e9kkbM8CSDEBdwBSliuQKH+EyP4VIHNnk8y",
	"lHCCJUEYFZypQy1RwYQgQqgzsCW6Jju0wZJwBUW4HgNzPWVSgTh4/2Z7C/KhoJyIBQ3g2mkuyYpwlJKc",
	"wawKwzK6JJJuiIKoIAnLU6F2o34yczrrUT2DWrBrocvueV18D0/OyZITse4iGvOJnmWKtmuarFGCcxfk",
	"7AqwMydbb00RhKBIWBG43pdnl6cvfz16PkV0iShcQaLQnMFRYJC9qJpsk4ySXP6vGq2nyFJecG3Y1kL/",
	"OXRYICoDPZdNBCYD6P1eUk7SyeFvPvfxFno3nUgqMzU2xPiqiTX1TaaTDzOJV0JNymiafJvQybuP08lR",
	"cn2uj9/e+lFybWHT5ovkhuSBMZfOURWaLmmimR98P1dcl/EUsJehlOOlRE+/QaIgSRCyMGrhLdFc8edy",
	"g/MZJzjFVxlBRxfHp6dIkg9SUeoNhbVwmlL1Oc4QzZeMb2BP0wrTsBBUSNi0wwtP1SUpershmTq6ooUy",
	"TwkXEueppUDYIpJrLBFLkpJzko4+pqIdjpNu5n/GWcEpkZjvkDMAmQHz1p2rieEIC434S0p4gEIKCxl9",
	"3PrbOXKWDO7bveEFTcModHoyEh4NGmguYrDCp4AKi4ci/g9YJuv6sqNEUAsML09PjtGVGuYiSZRAao6/",
	"MN/Anwc9bu19tZ60BpBCqzkAipx2b2D1i1cArR/a0IrJWdEH+peLl78i8Wle6ePbv9KwXXqXT7V3tRp8",
	"PiaxnLxcTg5/+6O14+FYpudt3PPk4zuFESs2s7iX78Zhot1uFyqOfLHqoccsX9JVyYEziIuyKBiXJMSF",
	"ciOEajatf7wiAhiQ4i3VRbiSsPo0/CIIvZRwxekARmeYbgJ8/EfG0UawxSZlCcJ5im6S/ynS2futRDcJ",
	"Ynm2m6OXersevmfqiWJLlOMNObjBWUlQgSkXSnoinCCCkzX8WLNwoSRPtQ2Er1ipjyNKPTdbLgnXorh/",
	"yjlSMotewEhkOAdRCIkyWVtQPsq1zJRiiRV9loksORGPp4hxT/53BoUfKQdjQD+g9qEfIP/X2z6ph/pz",
	"CrpSEFzgbLWAU4mF6MAVu+0EC4IEyQWV9IYYDiQ0WhgAGyUvWzFO5XojapwxiFIKooRWpLYAfzfqoc9n",
	"KkJuC5ZNLYbvCslWHBdrmiyuKEghiw2Ra5be4anWbNvEfCrQFSvz1ErOtZRgSedZns5eCcLRds0s11Wn",
	"93Fr1HFTKooM74IE3VYvHSpgHvnoTZjJUE2kducV3BwtDd6wWkPOcL4q8YoMVk8dvDSHCJ2PJWGlwWMR",
	"FVMwqqq9JvuuNLT3pp792+nFy/nT/3zy9JvZd++Cz5oWiANQRu7b21xWj9IwpMIB3RTROZlP0futXNwk",
	"i/dCPb0cZWmxuEnm6IQUREvPLHcnAtKcwl+a17csObAfkpGNgrI+nt2INlnkKXrEjCib7R6jAnNJkzLD",
	"XHNAjQTOBb84+qddAUY7ioHhlkAGrEIcf3wQkoynIRG7oj6tXCp+DHxacyNNfIq7wx43liPDZOq/dkis",
	"WZmlihObzdS66hucZUSOoysQjkCNbDCNWk86856yLkw/U5Mpta9+gBVq+zrGsNdXSWewt0fi8ZD3N/ia",
	"RAwB3cisDQH6zTMLU9H18iv2AN+4eNaNHDeJDFN64P03pJ4S9XJg6aE6mO6OHXLz6X0tZSEODw7Uuyw5",
	"Tq4Jn1Mil3PGVwcpSw7WcpMdgA42U3+fMVzK9UzvYHaTzJ487VXJDMdwpLpeqcwSdf3CzztFPq2NNiS+",
	"k/pB8GWtK5xcr7h6oBYJyxgPkkHGEpyRyE8r1ofoz9U3SunFm/AkknyQHcuXPAv8/WMIhvacEQBF4XNq",
	"5NGfqZCM706wxG2U6/wccVJwIoDLNhhmJeyu9efmCTZMuVMBDpkGXOIKm9WcCYBXRZStShJI/IdQjGOK",
	"oNQZUzqWAQ7yrPoAnWBJwoYcA6PIFBbg3ROEnpDTlnEmNLrgbEkzsrghXASNZWaaM/0dMt+FjZoc58LY",
	"pEL3d1n/Hr7AuF4IRhxz0sA1B9lKA1crS8J4JnKuTc9HN5hm+CojQ6wZDrK+KtTddniMbginS6pmPtOU",
	"BDjjGJi6mMzrzsFNmHYvFYSj3n5U625AaphRbKQ57D6UvC4Lp3lMXR1Hq8ZNbV8pOUR9WnsEtHnJmOvR",
	"mzXJq2ff97JNXVm2/lVJljjfaVeCu6D50spA9RDhudcMM+7jj5YaFiQHHXEAbNsmmGf12A6t4EdH7vfe",
	"Bw26qGfDiJ192/rlzSVIlJGXcazlcg+j5Rdprhzrcbh3U+awDQXNnJ2yAsjofbN3WUQb/C5s4aw3E5A6",
	"jq1e6nm5c6WxoQ3jASuPaGyYpFQyPtP7tpse4/Zuo3X9l5cVJ2+KOL3umoCFwT2XcyB7RprDF84FuQ/j",
	"aKHC19I8+tGOU1FeCUUCucx2TUexTxDARWsOqsnHE2dRziTiRJY8H+KLdRDUR5A2YN+FWdfelu6Yfb6L",
	"UIxc3vbBxCWXThK0SB80Yd6a4PZBTZy3N1ObDbRhBNE8ycqUCGtVwcl1zrYZSVcaPI7AMhpbbXBB0DSH",
	"0QlZEs5Jiiqx35lwji7BbAbWIPUfGr61Qd6+PYguI2aQLRaozMHrLRmimw1JKZYk22mwdJj1qegkJLs8",
	"ScBI7Ky8pXINP1dnc358lqcFo0FhJU5KnaTSxPf9KeeZJxcFrVMOt3FtgUousFJV2+LaETiWrQJP6ptn",
	"CGer2mcwYvp2cESehFcgeXI3K7zfXg8BF0aC5quMoKK8ymgCYgtWAvYvb/6hcWvvPTQQR21oCqDVx+/E",
	"HufO7wJxOhyM3RikrcnbNQEdoMelWAvwAZ+k0iai/Bzs6axQwy6fX4TwcbD7K+h3VHtR2PXb+Y/Hf//u",
	"6d/euXt1nGCPFILrlR7bj//zneNlMZbrvnNZdgICU56wtMnRlFwShwYIBL+8ubRb+P7dSHtQnnwieCly",
	"/beAlzncoqbYJrh+YCwjODfPkFZ+4bXspg4zoTZJOsFcLrG4yG/M82Emg0713VRPoeTW8dSxsrMUMLMb",
	"wndBOKq7UUchSyUyO5IICKQ6Jo24012TnWh74ZHRdNvbXeJMmP3amY/+iZI1E6QCI7XRb/7OYSnGleDr",
	"8NorfSnt4MMQx4gQRvj+B7LnO/ENXEgsS9EpEgv4pP1Ui2poBMv/6HmWzATm8+CpL7xPxh7rZSFjkXva",
	"F6XGgsrkieX+MYedpe8IaisDT/HsQ7LG+Yp4QdrHLCUDDItEjwWWWso1An625Gxj4xTBbRUIuqEklwss",
	"hPobi0Qfa1oCgrQuYLllivuJKRKkwBwbxovR28n/fjtByRpznEjCtRi9pFxI4JZUOCHDCEtJhLbCql81",
	"lWoLSceXZ+xMfR021DQOFAkzvtB2RMMidURIHd5ayrWOfJbE20NRZDY+18R1hDIW0KPXxxeP9cFZnu2c",
	"p6liSm8nJc8PKZHLQ7BkikO4n0O90qzaPqh9h++3cmZ/qeHwdqLTB/IUduqE05j9bkoh/cOUSvBELxWC",
	"oa/nT9BRPdvsB6yOf6yHHtWj1ME0gLoAHnRZ6blOTwBDXx9faIOhE4wbjgooFmpPA2iv+tKhv14iuj0x",
	"xgyj1Zu2uS1ZRvNb7iPXQ34wt9fD6D40TDX9EBoI6p+INE4mknqm6y6GtyJSat+DGdnJy2vPz6JwXD/t",
	"BWqHEnJ9RGpGa+GfXO0k6VW9Yis6AIyfuwtw5sCdkBPF3YFOX9Sr89Pgg+8c0zeHtQ/QMGLF4eVZDuvl",
	"B0BOFINBd6okJyyJNTZFIjJuEx37oswkLbKWnIxN7HIg/nWRBmMQzg2g4PrOOJlZmlMcW7GUHzO2ndcs",
	"9oLwG5oQhBMpEBbo5RmM3Gr513nHwBpbZuDYnRyCZB8IOIWdEaMvhfg8phtkf7enNxoBMDsda+hYk7Ud",
	"D2Jh11gYN0/tRsRLqcNnEyLEssyyHcKJAgEw0mZyU+QAQc9fZ26Mc0/uD90zWpdQyPl3oh7dhiVdOOFp",
	"CcsFTQlXd6TnSV1Gk2JJZpJuSM8WbIhN9DTwQU/ICNkUGZYkHLxgfgy4MV2PL9quaUb8e0sYWJS1GYsK",
	"7/Wv0sSm1mprIjKMhRfIUMtkpXpWLT25y5udibAxzDKMgdR+C81q4ArH6rkOCQ/3xlY6nBTneIvUhjIi",
	"STOYF2KwlXKelEKyjXnee+LRPwNaYDeEc5qG8dj+WBk8lGxKSQZehHqSCx3zP0cX1sxo0Izmq7EMp95P",
	"JFTppd2T/qBj/srwEF7AZCos0qDvxa7iXLNNbUhpOnrVgnB1NQt15kQubnBG08iDeaY/RfpTVH86ZNE/",
	"hYbVG/rt8amh1QHatx1oIiL0uBB9Vmaq4SJ/g6kbaiQCbRWfuKZ5CiGu+lGsXF0QkMjQit6At+v18UWn",
	"9mb2v6gC8kz0pb/4q/PnlglV0Y1mKOSsOhIAtpHW6BJfE4EKThIFjYQghbBGRV1sSZZd52xbRWrUgT1g",
	"ybtiSmnq2KRmUc3JMId0WmvUAwtj7rgI7XVVp1An29Isq+wbmutFvqR5Fc5fkJymM/vZzH52eHDQBe9q",
	"p0MS7LXUdrBmGXBHxwgB2GaU/frwiUcNr86fh3fS8RA1k0Vu/SQNUltHvqABTXbFcS4jFh9DGQnOK6Oy",
	"uWMYpUNgkVxzVq7WjaA143yuP3SEVjAaabnHVfZzv84FpMd4tiKwBECqDIi6khQgwpC83IAx2WMH6uPJ",
	"NGIzgm1pQ1HByQxXqoEe9q7HxBJEP5POBvFWIY+KgaYiPlbg30tiDWLGxG7jA61J7YpqM796c2bGke6a",
	"phRELAeonObt9SRDGEiDfJBIEInKAqUl7Ljg5IayUhhQWjeAoQ7FfegNRDHqo7kJCfqSp4gap4OJgVD/",
	"Nn6G2vvftIwZfm6PHwCRNjFaiDuxjiY0q1UdhObIM6loDW+Zsa0WnwKXrEDdFfpYxTuGaaMKVqk4JCC5",
	"uUQ4BvlQACdQKqYRxzXSG0EA7IqKHflYbsNF0AlZ4jLTj1KzFEZvVYpqf/C7GLYxN/SvTXnq/msl1N+f",
	"Zurj3HmlIHxR0C5n3kAlfpDPr3F418KkX1+1H3R2+ivCGVNjLU3Z+j2mvk0OwZQuPhnwqK1MQjKgfo2q",
	"xzitXuO493KZ4ZVw7NT2IEo4yd0gHwT6oZlYcZ06W2uAXBiW2vYT/cbLfA+y3oOs99nKerXelvQaNz9r",
	"4e9B6nuQ+h6kvgep70Hqe5D6ggLbvuJef7LTEHkvFtwMxWMcv2ZY5DKbiQgizsNjOHPNHgssFBln5Ea9",
	"VW4wbYNBs8DkcOu1qRvEsJ8vL8/QT88ugdfDP85JSjkYxfWyAm2gOojOFvnvc41BjihjGTs4+RQAFXLq",
	"8i7qOQa/oFwTytGGXSnSfVN5J8PZBR/CrikPLJb9Oh5OTfSMc5JpkNAlyglJI1l3lqRDBeU8itFg+4nk",
	"REc/vbw8Q4WWFivY9kdoBzFj2g62iCHsPvj++sxmujdcRSAZvTp/fqGEsnDSfrrL8YYmbmzEjzSThA8o",
	"fFEPOdGz2JGQ7Oj8Wlnu95m6NTg6+2kaPGJRcmsZDT9VAVXtuQliNsKl+2LpWhPCjcM1FVZqfzcQw89a",
	"6FcaHMRCaF/Q0McoxhrNZXfhyY1ZLoQpLmfsUGIdfTlAuKcn/eFFwenM4HfRs3UlZgMvcHKhg5EVNX83",
	"j2tXKEysRNlF5cQ0ypGS55Ym2i7grO72onZ68mmO3m/FIw3Ex4hx9F6wPEsf6ZkeG2VV3HFa5z3ETdx7",
	"0MJxG/AISikEFCMdC9Ono/oIZeKJfdIL4NxQFh2e/dZhzMlavav5KgTsNc5wvgJFAqcpqcqVQX51zJSA",
	"g5kdl2uCUidERE+hFDK2oVIxObETkmwQJEKD/cW82z0mizpQfVh6fh12DSXDNjj0lp/A30ecW/NILVK8",
	"gGjYMAhenZ9aCLSH1MldYQjpMGmSfv3dd0+/d7PD2BKdnJ6gR0a8YXVJkpPTk8d90Izjp0WygShaFRto",
	"iw5b2VF+nS5RXUMLkd9LnAmUbOUcXdBVrhShN5dKZa7S+qGyVZXaH8m1G73ie2fFX8avCBXZirGL6lFz",
	"9Jzm1yRFUDQIgNizfG9UXr1UfEtznR9elcdyahLopdXwOTouOdeZrbIds15/qMjlq/db+VW/WOtsznm8",
	"K/wZml/43NSZaqbmyYUkH2SkbBTtsW+BVFYVy8NAsjpA0dGUlIripPdmbMUCCYanVchONzjUphw4wLGG",
	"FauCWP2zqthJTIABTV8hkVPo1FXGnHIpSo8saZYaizLjJGy9QY/Ofzz+29+//f6xVn8164FBxpSqVU8T",
	"3WNiS8EC4c8Hlsp5LPWEhoVw86sgCSfhi25Zt+J2pREytHtr/gpuqkNzf3Yt546bFzeQxZ5xUmDeX7mg",
	"llvNiFDx6HsotW1Wq5f5AYdjMdZYrEka62nwM/xqzOTGTJtWNgFjvx1lBhhZXktPMw3CorF551IjtzPu",
	"biFcWr0DRxENqu+mdbw1vCS+yXi8w/E+8kk68nd6bdKv60wzpU1pk9XbScJS8nbSbTy+IyIP5RQNuri7",
	"QYJ+O+QALIjWWPDQIB7Qr3n9V6LB7X22TuLlK5oNeviwYm5NlulU6FPz6XtZSJkFmz2AOFzVB4MUN+17",
	"ubx8Hi6hVJRA5cG9jofO2dF5N0wGcS4oc2WMlQSVRcI2bV8G7ypC0TLVLzO2HUjiWvixNpb0x4xtQYXt",
	"NNZU1zuNIdi0YreR+xxOa+PMnq3XSouPmTGL7PPQDSDMAU/wF/I69ryDY5/A4JUAXEPGdv8zpL7TmZUh",
	"xpZSkicaa8KK+Vv10duJcf+NAmowoP5E06ruxGTSXxxTXx0qAHXgRxUS379kYGVRXtyu2OG5naev6mGk",
	"0GxdwRviK/pPsufjrZefNu6/C1EB2/blHudElNkwEW1QUcR/6xJ+n6REXz8JtUhzjk5zSXiOMwgbgALi",
	"ezVr+jNKrDGaJovYobVC36yZOvS8nEi+6ytBY9z5zcAEczKTzAUz+ZGDVURItZdXgszV/2h7aGrrowU8",
	"p7F6K16asIGJPYYPqWmgKZaL710Mw9D8cJbR7cb0OQDknomR2ouz1Y61Bjv8OvyybTOvDsKK+z//y3xR",
	"8fzXx+MiaiJq4bG2VHbatjsOshcwYtze/UZx9zKTsYIDIUeJN9z1EThz9WB+Nfm78OkH4Gzj1K2eFLGa",
	"bb7xdEMkhviOul+SYy4e2JLCB5y2HP+JTYgCPStqU/n+osqAc3m32bqRoRyoFOuQpWaIfakU64YtwQyO",
	"KzKfi2UpVhkm1ivYhXUPxEYAnqTjzTkwbLAJp6sAtakGnJebK3iXsWy2p6gKURvubE39r85P3drUUIyy",
	"YIaKjMVEFzRyR9RlrQUyNJRSkXDilmMMVki6KqV2WspdQROcZTtd/yHDasUMOvhwiR6R+Wo+RVdEbgnJ",
	"0XcQjfa3J0/sRh/H+ulqE07QFdQ8BBhbFLStaN3edJWuwJRKZMLCAWSiquU5U7IVV1IMMXXLG1VhvXC4",
	"doBxOIC2V1F3j+p1KW7gdwwxhzriTLUVU1ekLRgI/cOzqNHMViTptpCFa3WZoXEJoFWPbdrakAOPxlkC",
	"rlX/i1NTsSJ66sHsrrFy35Nhp38X3OKKCkk42Ex1oa0oh6/rfVVx8Wowb+Fa0EEXjtNptBE2FcihTRfz",
	"YvC1EcRwoTrSguTpDHyFJpDfw/+udJogUb86f263AHHQW3KFCrwiTk/jdoXZHqsPCDmJ7LJuWPmi4rI6",
	"hWcntBUXxqOCsCKr6lNTBa1KstDLTx02SDaYZginKYd+gOOE5zoTpmvXNTr4OTB+8TjF27KMbavyK1WI",
	"sK1jJw5RO19livZJVxl3zPfbaxGrNveV0I/gG3KF/kF26IJIlLKkBM3a9Mwz3dPdboeJHVwH3YTbpam1",
	"e3HQvgM21iIJbu3RL2/+8djb4D5b85ty9W7NSAXmnVLvF4Q22JikDnooWEaT3bAFgOcKnbiz9jlFwekN",
	"TnZIT1ffTSPLzPbUTEmRsR18wfgK53U6R5bpPpalIGKKOAGITUFEUFJIxgQRqCBcQMgt5HuEzR46rl0d",
	"rItqLDHY73W+3WnFAxoQRFXeB11akqo0izbZOKQ4jhY89+UwqvfSfdqEn+AcDEXmrxGnX4AZjCfkSOLP",
	"RaCdhShwQmZ1rVFbNdrpRBg/SqsVTLOCWNurzJZyizmJdN0oc/p76XV1NdgPEit69er05DF0yYdYM5Ou",
	"YTZV98dnHNl1NHGLtW6CC0O8R9rCHWjK02QtbtmJ9HtrwvThSeFGVIj4DKqjRrugHdnGZ4ED+2hfb6P6",
	"Es7y1gVoxJEPt1H5DLWHcBOJtKz8E1Vp1lC90mpz2pHWhbs5y8kUeUE9CyXuN/92hQVN5uhXlpMq0VGt",
	"Yniz/ligRzkoMggXhZja/Bb1j8eWw+MckpbX+AYK3nIiRZWOdhhcNAwzcWuGLAnfgLnOiJs1S27cbYND",
	"65RMJRiX4OXQ2TViTYtKYfMEPVPo25vN/yBJSCGFplbLdvwntNsA2CETD9dzAlP01nuF6LuazGqzGKQe",
	"mXTaphTeExEXLKXb08+wmkAXG0uDddgulcaOpUFEV+KriXuLRdvN6Daf+ixVgzpYMAg8/bNR36tKzG5C",
	"HSTH1yUq7Sb9etAsxFJ6d9VZGi96JXqsNpXoCdSj8UTJFNT8WXER/VPnVT2oTQ9q04Pa9KA2PahND2rT",
	"g9r0oDY9qE1/ebXJi6Vpp7Z4WkQnnvkS1LsehWy052lIKOKAFnN1tv1Du8JQ/n2oSeAw4A90kF8Q2eoq",
	"fyGhu3rluBqWb/8r2ZoaCvOeguR7pK33VVDrSTUPhlePT3wf0xTUki0Ay7m9XoDf/uJshFSjc1ZPI9jR",
	"4ej+fMOOOCZs70IyvlfbLSEZH91zi6Xh3KjOxKlPkdbhRM1UReQsoDshdEswj2iotA/AOxoc9R1vXLrJ",
	"qyLFkjRLEUTRqPPzKmBESF4mWqoo1QB1+tfH0eZ8NVsIVl25fWUFJy0rsoL59XUtGPcmi5vZWmOn/nkC",
	"u3dwtBv8A+/wtW5PQM5qfCDpQG5gWxvowoSt8mpKlCtoPn/owvfQhe+z78IXKioaimVGDSwfWVTtlVJh",
	"DFH0cYlwlVND/L10e3v67w/m3JcBDCwIXVUa8XQ9b5BTadQpxGrfkiofA8z9CeHARdxclV1BEBamaBlU",
	"Jb0wVrvv5k/nTwHXW7VLmVwTvqXQLFibwNtlhKeRaf+uvvnt/Mfj77/5/m/vxmYN7RM53CynpNOE48nl",
	"IfNgZUhrXLMZMCrjIZyy6RXITPvrCNaiW7WHVhHBftweSiRVsz4/7SCmx3WXoIKfTJpbMP+yu4hTfCB1",
	"IjeHx2VW8Z4fp5PfSxJKinIoxkvn+G/1eUAnbVyWnrU62NQBkLNp9+I64R1QgWHAzimivSbJdSyhRX/s",
	"Zo5Vtm7HfrLENCs5QYmaChl2Eyo3RpLr0D2rUXCesLWDcxZojvVM/RltiBB4RfYuzPXazbKJvqJN/RoO",
	"YncWXKh5QxGAD87DaU7SV7LQuTF3d/2ZpgUniS7jpotL3b7cIPoBJ9dbzNVLtymwpFc0o3IHfiZU9/Y8",
	"9grV3TppdWDhviZcq8p9bv/S48+uAOPHOHqNq+kZO39nbcGbJle479KCd1SrrwNqQ8rddQJuiAxY8U62",
	"9Dod9lCo4hfDE6e62E1Xzmj0QCNB4uVa9nCpIpbwGKiTPYae3T0EKfqss9Oz93B/WqoO9z0+i7SJjkD8",
	"Fjc2hv6LRhfsPSS7P58FhA5/C/iNZQMjKCDAB3o1iKSRrD6qmlszaTgwP0hAI/dUDKc9Df+gJ6cIpmOP",
	"P5GbjR1qbsCVhCtuUUWsCGVft+9lau93Gszh7sC0wdj6hmTZP3K2zV8WJD890UnNPS3N+8c0E0lNn1P/",
	"C4PwIJliQYxL+vXxhTbJQV7p6cnZ/nXJnJ5UL8++Eq4JzbMAPusKy7zCMlm79XEGrddKYf9KtCsuVuva",
	"FNHn2laiZGA1yVrKQiBAVW0MenH0z8qWWzAup6jAcg0/gY7oWHNqXHdLBk8j+fUpIwIiJYzVEz6L73dM",
	"06hGJn5d1v7Mu9NhLgUPhUSd7P5xund77FDllXgBAtckZq6NeaEDEOtnjD053pADp8Lq1NSNJThZ6wBn",
	"yAhuhzmZrdUm6FYpJHugtM+jvTe2fno87XWkW/h0FncY1Bmn44I5kSXwdxRe27WY5m13QmVYtT10DJdz",
	"OkjpNjvcL49j1m8Ta6qjlWovxBJngoTb7rg7jrQPv4xcd1+qwK3qg3WFgzSIWNf1uRN+e2KLBN09Kk/v",
	"i+d27jlcbE4UGd4Nas3n8Z8m2zITofqp1Vb/9sahTV3lDVBqe2msbYOESMcqYfbeHWPfRewQ6a2P6cWx",
	"Wg4MT3/16v8E0b+Xu1YoKxXo5ZlXZGa45OqVitobV391ZvnskTS82QFeQX2rOGf5bsNKsdARwr0XbFl6",
	"pJqYCeawgY240b4M2C0O9lLTdUXkmpVSYbTNa9JeYMt4u1muGz88QhQ1Baes5/bcjULuhKgfiX53tOHN",
	"e4fkob1Ld7fP30yV+XfBmHQqrDt/z91CKPnCJuRFg+Zt50qMRNUfwlDrL28ua6baJqgq188psY9FOzwx",
	"FrE9RsvRdNCJTvEw3VvdWVe8uHDkWojZp6IVOn5S097bSc5yU817j6p3g3TVcd7MLl9yxMKjE9VuyM6I",
	"dFpErBPknEKY+t+mca/t1axbS3SW5veUrqCL2A8HMCUvPf+/7ousO4TaLiCOBN2hqhxHFRQrJQ/QD9EG",
	"F/bzNuW8MJRjC0uZ5KrUSjIj19H3IdwLmSJRJmvbjFWbcKdeg75GISxPTmvFOthmh3RpW0J7wO5qxeqF",
	"JNtSjaMro9ZVHuOqVLAWquEctU0yz3Zz1EzInGpBUuNEtvMrm7b06gTnCiim3qmOaVKgPNJZQnZFpXxZ",
	"pKF1r9k5egZCK4Bf+ARTl+QS1tIrlHJNcsl3Y9Cig1srSUGdwEpZKV0C/klTINo8QyAZ179BwS7DEKGF",
	"nxI27Ib0X9TvfhYPQKsRYzPuFYsVfPUZQMSqcpr28wT3pmtSCVpUqmKvOSGpMG/kcbPorMtRUrXyBnog",
	"LksOUp6BVUjfrpip4QK6VWcrf6sZs2Smcloq2uK3URru4LXDCTtjps1ot8LmVaKEnEhRC6iSWaZseg0l",
	"2K2gx4lgJU+IkYQeicd6horgPPBdaHFJMrSh8GLrythqFslpokDTyuO7bXlWV4hybjt+t7rtVKvNh27B",
	"7NgK3k6AW+tmj8cDi785BvXgkx6yr3Vw2gGpRfWTQ1I3lSRQ+NwSv//s/1e0sK5i3tBCLPGYth9LB/rj",
	"4cHBdrudb7+ZM746uDw/uElmitXNIHj+4D9sZd5xF+7EW5caXCE8Zy2QjDKaWl6haUKxX2iLW8sa7qNc",
	"h0KbFwjnEH+o33qzk0e55gLA63UEfcmJeDwFju4SZj0onEwUxnh9CylR2gKuEq90ZnBtSmnce6XD3K7g",
	"vYvgIYy1MKgFjfkgnI/WXR8UgSQQFLPW6RjJ+PL3d+tVd+YOQit61jigzji9wZIcnZ1CUNoQH3Chh6Cj",
	"s1METtQuaFj5L4gUneFx0bBrGNXMgWt9pbjHQlS5IO0SIjQHiSyRumRxJHeJ1N7Fbiatz+IVOw5DtuMu",
	"IG941FXoTOMhN3F/sO4FRPhYcTjEErsO/2j1jwxkVno2iBFtdsclrtmWP5Gx9uebaLaQG4YdZ5uRfpnR",
	"Ytx2T+0dOBHUDTBFgRJuGxa7nXZ4MBDZktkS+yZkGYraTA4na5Jl7L8kL4W8ylgyT8nNZDrRaaSTS/Xn",
	"HzKWIEnwRqFmyWGQkQn8YS2/aj0c/PTGKOzw70pTxnnql6vXKt+bb47R6+OZIiycsXylzbPaOPPta+iP",
	"JVnC3B7bBzZgwU1C0eM0yNQpMpoQQ9PmpEcFTtZk9vX8SeuQSvDB8DMIP2asOHh+evzs14tnasxcftDv",
	"eSuk2tUlbK1hSB7SURoaKydP5mphzelyXNDJ4eSb+RPYS4HlGpDywJzPeRUPasZasHgSoHBBXqf2VVz1",
	"NFXsgQkn21WYBLiq3vUPLN1VTRo0UTu5UgfvhaYubV3osz1059J9/PjRMV3D6b5+8mTU4gESaIjZ/1DQ",
	"/nbktAPsKREuG9jCDzi1GqXey9M/by+v8trvojfzzZ+3mR8Zv6JpSnRrClFuNpjv+pC5ZXn9OK0oZsVZ",
	"WYiDP+D/T08+Bkjo4A/9/6cnH9V5VqFKb+dEckpujO45gKR+IkGKKpyGu7+F+/+hn9RWTfo6VX9XbKDm",
	"y+YkE/fh0S9TiwZq8aCtWOsTh5cQ9a/D13j3KejWQ4oBl9KFGs4bIQ7Ih2SN81XthNRZkzY7Mcxin5lB",
	"DWNEOHu8aujYRhY7T0ca/H2w4t5l74Ab77l+nEPcH/eOKAt97NtDyf0wYgyiFrqP1AwkSG2HUXP9a+b0",
	"Qg1jq+lAZf1+wQ6/rrOxlqv8hqQB+UHPHOleex+oO6hx7j2j77CGol8cCg9tBr0X0nqWmYjcauqtVQnl",
	"DmO3CpObeuzawsB5Ck3wTC661TGWGdtG8dbrNnqf2Fqv84lQs9ls7stFRq8h7P5oN4Mg7btDPpjOT8ja",
	"Fwvbvd/vERWbi90BPu7X5T+aWvaFImozFH4UupZi3ZA/ex/1FsIae6TbNBwKBoMChdzkfx3uFnKjBXA0",
	"0srsvrC0p3NaHF2/KJyJtsUbgzVCMj5ObYGaVOK2Sktf4a77wIvuNe/5Ue0p5fXFsax90GAMYpqaNWTm",
	"x/b2IKctJSKihW5Kp7KPj5IDSvXcB1b2LnvPiNlffeWLw83hWNCDkcYHIw7+qIrBfdS/pbNGZ/KY4a/k",
	"7fhcEOfWVPHeXRsP64/ttz/rTye3xIKRsbWOD72KTzNhdVc7tKI3JEcGLHskZTTOpgs//ply3H7m94ZJ",
	"Ucv1PdcdiPXttOyeaSCjmMHVLVR4C6tutVXp1UG1a9rqQGZR+QF+GjG/V402MmujrGKHvbiPSv/wSzb6",
	"pnsYCK/IAIt6Df75vcPfWc5svHvNuiDlKFN7+PE0G3AAFXb2wXeOf+KeHsbGMqY07Jfs4buLtw2g4tzR",
	"Pi9YizY8yQsOIMgM5+nM1vWdWSvEA9FEVGgnQ0wyZOEGWvVpMHDBDTKgUI/AZhz4Rc5EPVk19tX5c6f1",
	"gK136K4L4dIZ23rSuIPBAdK2ESFuwDZggn0Y7ovOzboQTnt8+onE3saq5qjO4l8YW3ARDpnThASRu+cX",
	"FY9gNE0e+MNfiD/8FRjDn6MF34Ey0uAIn4ITcF2w9IEHRHhATf8GUi7hm+SjK6/iqQni9Cynsd4S92U1",
	"7WvXcd+G055eGiG6vP6M6fKCSL8VjsYEe4VdJNlJifMtybLZdc62+YHOWHI9l3Wdlk5DlR0IiWftJ+Al",
	"/Ow/ADbNeXKPODCgetiXait6fXyBTk/OAqXLPmNTUYOV3j0nVUivxMmDyngbRd1YtTUDYNuF09wCpO7q",
	"9oxVvmEzc9Ntn9vAf5omlZG8LyDzdd0Yw80mTFhqywwEDGBefYzbXdJlqIdybF2309ot1jxCVelNlBJO",
	"b4hJvdJ9jFNbdBlKqujaFjrvNZRyy+0N6i6ItjQhwisl9kqUYdlxIJaShVvA/FanMk1AYM9bXHcw0GfU",
	"J6sWG7aluk3dyDsNthWxjUx1NGspCJ/hlcni9frOuh1PKxd/wckNZaXIdogIiXXzytRUEIstafpgOz1F",
	"vCaXBWdAX4zrPPANvrafB685ThF1S9fxwNLVW0zCq6H4ngV1H9NxCJIjVuDfS5sC6nXvrrKpN5jqrH5I",
	"yfH6KtqIIJynKMFZdoWTa63uBUFPdYCQ0MUL9JqmLaq5XQNpBxHUlD426AXqnPKLn1++en5SqYumSPKN",
	"6YSdcCbETFBZ73bJ+Mp0FggCsur8MBiQz3JFJGmdCBsvfBcorOG2Am9UOtF9ylqVTtCLMpO0yKKLOOqz",
	"poadQqdmQrgb1OVfGM2h6Kg6ysYu1TCwh0AX7rwyCpQ6Z+krYZKelGyRk0TakhOvzp/r+zf/hq7ttvJX",
	"SkXCbqCgl6Fi4HUmhd4B6FcKRAWGqvqQKJ2ndXfbOTp/dvzyxYtnv548O1GQqKpRueJvJy3a9nJa/NmT",
	"JsHZu4ZQphoTXhz9E46ryLG8Emob0GNY057GkULSDf0XqSjpK4HIh4JwKF1wB6eD/kNrncI3Kv0CGK+p",
	"1OhVQ7Ep3ubabONl8kHaDtANUxPhc3Rkpqoa2XtlJupu9gUWQie/49y1U4GdweHk9YtfK7w15E1tDt4M",
	"GXcbA6mVYIiZQScxmm16jKx9mst6XehxJfE1GNOYYv+stM1qbWakrT2yKrGSConeAON0RXP1szkLFWbS",
	"KUpYmaU6wx9hKRWnjtyvu/m9rtipRKeL21Td/HWWI/aaOKtjNNtUh56Pju5jPa3HaDrT5QD1n2eWT+Cr",
	"jJgmZG8ntvYtEUratXLl20m7omnFMqE108+Xl2cX6Ao6jb06f+7U24Hz6wuvZiw5hR5nyw4BxRYVxBkn",
	"ON3p/symp5uOZ7BNkZ3SOHr+KaK6Eg83aUKNcVBOAL78f//n/wpU695VnZReSXuhQTkZkxb1zZOvO5TZ",
	"D7PtdjtbMr6ZlTwj+i31tdtwz89w/Z6QAKLb25OcVF39urEsMBo0IvKhoJwgsWZcZjuEl4AWgNomxgRK",
	"xUi6stZqTsW1ekYzgq/FmGY9R0jRSkbQ0cXxqc2zH5cw70/4s99yW08LPFYHugFlVvWn3b7O07qNoBBU",
	"yHAneZqjMk8JFxIb3F8Ts23NupKk5NEG8OF+YlVNJro0dAQfelSpFJuqEFHuNTcMCuxwweQDTmzVYE4S",
	"0lD5hjb6tm3z+pzzP7IyTxumFDCd9GWZ1M27K9tCs+R6PKzwsqtMuUZYUct3btEwmiOWBwZXBZsV7ysU",
	"3tTU9CxPZ9CAsCxAj3KbEiwR1u38mgXNzGcaXaid1DQ0ahkxPk26QGOVT2Qwbq1ameem/qxbGYwMqVD0",
	"yzIoAhV0JBAECGAI6p9q5E58nLYJr7pIc6Ppoy7hFUa8e8e5T45uf1FMG4pjNC0+b7/F668fPBcPnou7",
	"91y4Bds/GXs9ShQhZSRdkQ3J7ytl4Si57mSu3wa8NH++SxY+CLEyt858NzMrMI9fpSmmLEie2iz5oNyM",
	"tIk229l24i3FFecpWhFZG0lenZ8qtKhCf9xKt6BRYFFXKraqsg7Y98xbdr7Wwt0ur7NSrEl6qwT90arp",
	"wN7TLYPxv7mxeEyL9agDsD2J7yw7/Dzcej3btA60wztw13XWjvvrWl8rI+nnbHkN9K5wPIaHfzEXancn",
	"j2DGZHeUQrjsYRiuPd7WocaqB3dqC1K1q+jwM3d0tbbu+/AOv3g/ZbclsxnA4wbWNJ7ZkL2zLUw/vdO6",
	"CC0xLi4vH3OCpa3qFyiP0CpD+F2g0q1+iH9lEh1lGdua6Z5+E1LlNRU8yyWVO3TJGHqO+YrAgK+/DzAc",
	"xtALnO/s3YiQbK/PvI912BhUXXm/VXlFfRCG573JxTEnR+M9MIpC7eyYRGvRjvJ9dK7jfjpgOZouQJUN",
	"aMUnxq5dd2g1WrDT5ge8DYVm8hUHrzx0tXT/+kxPNh+yp6jQV0kkYZ0Oqo4zzklWqZHNqW+K2HHtDutj",
	"sJyoV3/DOJgqbElYtwevGHCej4M4TIC2L0p4vD9jK9QrI6yK8mpD2x4qqygzVzPhrFyt0evjiybl3xQu",
	"5dtXPx5yqjiL/Qqufo3zNFPPet1uuM6rUW+bWwJPiyVMyQElQaw0FfKqUNdIpSeliZ/brfVY06Dshq7G",
	"Udfhc8pwxMITb2dcs4EOXcFg+5cE/eZJ8NUwAAnwfgdYHXy+oslOA53bnATujxXavZvtEK7SOfTPNqig",
	"suI1zRL6ZtyIjjUWxsqgFGFwhgtNgssyiyB3GEOAkdzf89NhbrAu5qn1MdfRKhCE4XBvW487GjswhFWH",
	"gxciQQdt1/St1l1Uzc1DthK+KyRbcVysje7OcZ6yjWm3U5lIrL5t3w0S1+yspiGNt7cSRnt3W3c4Haz7",
	"+datDk2w0XazO5oE0MKOABY3ZPvdunwL5d56A1rRDeZ9TXsMU4q+5ZpQbvttWRBpc0+iveq9e5cfRoNE",
	"L63HheIxHI3k5XI5CGEb+omDD++GSwd35DZQDA0Y1BeQVXke9316jV+7X5dO16c29mgm8pAw6b3qGjAC",
	"pVoV1+9s7jQdNY9L9Yy8Pr6IsvSQHKUX0D6be3KT2UVg03qlTrfZ0/tdeYym/1mSnT2VwYLq7sLkZ2Sy",
	"TgqMSvgj2l+YDTuZqYNr9X/Cwk5tcr5zar6Lgk+frmHA0JAeuNWjG0zhje1/t4Le718ZMhjVwOufiKyU",
	"B41gjf5lbiiFTa+HWIoI3wQp19RrTdEjM4Skj7vLOv1ELAKT1IvseUDjT4DGd//0hO/znPx+3zJebGFR",
	"DGzKMRiB21ShuL7Vy/zCFM2yyHVjtrBp93hNkusHw+6DYffBsPvpS1nZaimV4dYtS+QXT9IeQy9CHrTo",
	"sKm3stl0sYo/5Ado9JRhunHEw6bMp5OCTp2R0PHiHkq06s6zTolWVyYtI+1BvzCtuvfKV0TaAjmVmdQE",
	"TxgDvlsvax6+9D5p5gQiF+rip2HRwhQ+HRkSWiHb+GKlMHTXbyM4sYEXFRTd6rz3Jpa9bqyGbj6BuaBd",
	"lBS2sbv/qqTNde6qLOmYNb/gytxwknD50hh7vv+ShH9dyqmK3dE0cR6zT1HQ7/XZpyCdxpJfbDW/Ty4U",
	"DaM6d5U7eKn+FHL7M94pV96/14fKXejTPVXuql/8Y1X4dxUiHDUMvA0a3esu0IcHBxlLcLZmQh7+55O/",
	"P5ko7DBTNBFUh2nMtC84RRuWkqwRqtisLTFpo7nd18B5qmMEwjl0dOya4EyuUbImyXU9Tv9V//Hju4//",
	"PwAA///xYScd0k8BAA==",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	for rawPath, rawFunc := range externalRef0.PathToRawSpec(path.Join(path.Dir(pathToFile), "./common.yaml")) {
		if _, ok := res[rawPath]; ok {
			// it is not possible to compare functions in golang, so always overwrite the old value
		}
		res[rawPath] = rawFunc
	}
	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
