package main

import "fmt"

type laptop struct{
    merek string;
    seri string;
    tahunrilis int;
}

func main(){
 var kumpulan []laptop

    kumpulan = []laptop{
     laptop{
         merek: "Asus",
         seri : "ROG STRIX GL503VD",
         tahunrilis: 2018,
     },
     laptop{
         merek: "Acer",
         seri : "Predator Orion 5000",
         tahunrilis: 2019,
     },
     laptop{
         merek: "Lenovo",
         seri : "Legion Y740",
         tahunrilis: 2019,
     },
     laptop{
         merek: "HP",
         seri : "Omen 15",
         tahunrilis: 2020,
     },
     laptop{
         merek: "Asus",
         seri: "TUF Gaming FX505DD",
         tahunrilis: 2020,
     },
     laptop{
         merek: "Apple",
         seri : "Macbook Pro",
         tahunrilis: 2017,
     },
     laptop{
         merek: "HP",
         seri : "ZBook 15u",
         tahunrilis: 2019,
     },
     laptop{
         merek: "Asus",
         seri : "Vivabook Ultra A412",
         tahunrilis: 2019,
     }, 
     laptop{
         merek: "Lenovo",
         seri : "Ideapad 120S",
         tahunrilis: 2018,
     }, 
 }

 fmt.Println("-------Data Laptop--------")
 fmt.Println(kumpulan[0])
 fmt.Println(kumpulan[1])
 fmt.Println(kumpulan[2])
 fmt.Println(kumpulan[3])
 fmt.Println(kumpulan[4])
 fmt.Println(kumpulan[5])
 fmt.Println(kumpulan[6])
 fmt.Println(kumpulan[7])
 fmt.Println(kumpulan[8])
 fmt.Println(kumpulan[9])
 
}

 
 
