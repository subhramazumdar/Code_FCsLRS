/*Anonymous endorsement of Transaction in Hyperledger Fabric : Fabric's Constant-sized Linkable Ring Signature*/
/*Written by Subhra Mazumdar, subhra.mazumdar1993@gmail.com , Indian Statistical Institute Kolkata, March,2018*/
/*Work done as a part of Masters dissertation*/
package main

import (
	
    //    "crypto/rand"
//	"io"
	"math/big"
        "golang.org/x/crypto/sha3"
	"bytes"
        "fmt"
        "time"
        "os"
)

func verify_signature(f *os.File,mesg string,b1 *big.Int,b2 *big.Int,b3 *big.Int,b4 *big.Int,b5 *big.Int,u1 *big.Int,u2 *big.Int,u3 *big.Int,u4 *big.Int,u5 *big.Int,u6 *big.Int,u7 *big.Int,u8 *big.Int,u9 *big.Int,v *big.Int,T []*big.Int,tag *big.Int,N *big.Int) {
       //value := N //new(big.Int).Div(N,big.NewInt(4))
       //value = new(big.Int).Sub(value,big.NewInt(1))	
       str := make([]string,9)	
       
       str[0] = u1.String()
       str[1] = u2.String()
       str[2] = u3.String()
       str[3] = u4.String()
       str[4] = u5.String()
       str[5] = u6.String()
       str[6] = u7.String()
       str[7] = u8.String()
       str[8] = u9.String()
       var b bytes.Buffer
       
	b.WriteString(mesg)
        start = time.Now()
	for i := range str {
		b.WriteString(str[i])
	}
// A hash needs to be 64 bytes long to have 256-bit collision resistance.
	hash := make([]byte, 16)
// Compute a 64-byte hash of buf and put it in h.
	sha3.ShakeSum128(hash, b.Bytes())	
//	fmt.Printf("%x\n", hash)
        buffer := bytes.NewBuffer(hash[:])
	c := new(big.Int).SetBytes(buffer.Bytes())
        val :=  new(big.Int).Exp(g1,b1,N)
        //c = c.Mod(c,value)
        temp := new(big.Int).Exp(T[0],c,N)
        temp = temp.Mul(temp,u1)
        temp = temp.Mod(temp,N)
//	test := new(big.Int).Mul(val,temp)
  //      test = test.Mod(test,N)
	if temp.Cmp(val) != 0 {

		fmt.Println("Error")
	}
	
	sum := new(big.Int).Add(b1,b2)
	val = new(big.Int).Exp(z,sum,N)
	val11 := new(big.Int).Exp(h,b1,N)
	val = val.Mul(val,val11)
	val = val.Mod(val,N)

	temp = new(big.Int).Mul(u2,u3)
	temp = temp.Mod(temp,N)
	temp1 := new(big.Int).Exp(T[1],c,N)
	temp = temp.Mul(temp,temp1)
	temp = temp.Mod(temp,N)
	if temp.Cmp(val) != 0 {

		fmt.Println("Error")
	}
	
	

	val = new(big.Int).Exp(g1,b3,N)
	temp = new(big.Int).Exp(T[0],b2,N)	
	if temp.Cmp(val) != 0 {

		fmt.Println("Error")
	}
	
	val = new(big.Int).Exp(g1,b5,N)
	temp = new(big.Int).Exp(T[0],b4,N)	
	if temp.Cmp(val) != 0 {

		fmt.Println("Error")
	}
	

	val = new(big.Int).Exp(g1,b4,N)
	val11 = new(big.Int).Exp(s,b1,N)
	val = val.Mul(val,val11)
	val = val.Mod(val,N)

	temp = new(big.Int).Exp(T[2],c,N)	
	temp1 = new(big.Int).Mul(u4,u5)
	temp1 =temp1.Mod(temp1,N)
	temp = temp.Mul(temp,temp1)
	temp=temp.Mod(temp,N)
	if temp.Cmp(val) != 0 {

		fmt.Println("Error")
	}
		 

	temp1 = new(big.Int).Exp(v,c,N)
	temp = new(big.Int).Mul(u6,temp1)
	temp = temp.Mod(temp,N)
	temp1 = new(big.Int).Exp(y,b3,N)
	temp = temp.Mul(temp,temp1)
	temp = temp.Mod(temp,N)
	
	val = new(big.Int).Exp(T[3],b2,N)
	if temp.Cmp(val) != 0 {

		fmt.Println("Error")
	}
	
	temp1 = new(big.Int).Exp(t,b5,N)
	temp = new(big.Int).Exp(g1,b2,N)
	temp = temp.Mul(temp,temp1)
	temp = temp.Mod(temp,N)
	temp = temp.Mul(temp,u7)
	temp = temp.Mod(temp,N)

	val = new(big.Int).Exp(T[4],b4,N)
	val1 := new(big.Int).Exp(g1,c,N)
	val = val.Mul(val,val1)
	val = val.Mod(val,N)
	val = val.Mul(val,u9)
	val = val.Mod(val,N)
	if temp.Cmp(val) != 0 {

		fmt.Println("Error")
	}
	
	
	val1 = new(big.Int).Mul(bigTwo,c)
	val = new(big.Int).Exp(tag,val1,N)
	val1 = new(big.Int).Mul(bigTwo,b1)
	val2 := new(big.Int).Exp(s,val1,N)
	val = val.Mul(val,val2)
	val =val.Mod(val,N)
	val1 = new(big.Int).Exp(t,b1,N)
	val = val.Mul(val,val1)
	val= val.Mod(val,N)
	
	

	temp1 = new(big.Int).Mul(bigTwo,c)
	temp = temp.Exp(T[2],temp1,N)
	temp1 = new(big.Int).Exp(T[4],c,N)
	temp = temp.Mul(temp,temp1)
	temp = temp.Mod(temp,N)
	temp2 := new(big.Int).Exp(u4,bigTwo,N)
	temp = temp.Mul(temp,temp2)
	temp = temp.Mod(temp,N)
	temp = temp.Mul(temp,u8)
	temp = temp.Mod(temp,N)
	if temp.Cmp(val) != 0 {

		fmt.Println("Error")
	}
	         elapsed := time.Since(start)
        
	f.WriteString(" Verifying took  " )
        f.WriteString(elapsed.String())	
        f.WriteString("\n")						 
		
}

