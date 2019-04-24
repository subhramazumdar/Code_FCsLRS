/*Anonymous endorsement of Transaction in Hyperledger Fabric : Fabric's Constant-sized Linkable Ring Signature*/
package main

import (
	
        "crypto/rand"
	"io"
	"math/big"
        "golang.org/x/crypto/sha3"
	"bytes"
        "fmt"
        "os"
	"time"
)

func proof_signer(f *os.File,mesg string,random io.Reader,T []*big.Int,w *big.Int,v *big.Int,r *big.Int,pub_key *big.Int,pri_key1 *big.Int,pri_key2 *big.Int,tag *big.Int,N *big.Int) {
       value := new(big.Int).Div(N,big.NewInt(4))
       value = new(big.Int).Sub(value,big.NewInt(1))	
       a1,_ := rand.Int(random,value) 	
       a2,_ := rand.Int(random,value)
       a3,_ :=rand.Int(random,value)	
       u1 := new(big.Int).Exp(g1,a1,N)	
       sum := new(big.Int).Add(a1,a2)
       u2 := new(big.Int).Exp(z,sum,N)
       u3 := new(big.Int).Exp(h,a1,N)
       u4 := new(big.Int).Exp(s,a1,N)
       u5 := new(big.Int).Exp(g1,a3,N)
       u6 := new(big.Int).Exp(w,a2,N)
       u8 := new(big.Int).Exp(t,a1,N)
       mul := new(big.Int).Mul(pri_key1,bigTwo)
       mul = new(big.Int).Mul(mul,a3)
       u7 := new(big.Int).Exp(g1,mul,N)
       u9 := new(big.Int).Exp(g1,a2,N)
       //message,_ := rand.Prime(random,512)
       
       str := make([]string,9)	
      // mesg := message.String()
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
	for i := range str {
		b.WriteString(str[i])
	}
// A hash needs to be 64 bytes long to have 256-bit collision resistance.
	hash := make([]byte, 16)
// Compute a 64-byte hash of buf and put it in h.
	sha3.ShakeSum128(hash, b.Bytes())	
	fmt.Printf("%x\n", hash)
        buffer := bytes.NewBuffer(hash[:])
	c := new(big.Int).SetBytes(buffer.Bytes())
        b1 := new(big.Int).Mul(c,r)
        b1 = b1.Add(a1,b1)
        //b1 = b1.Mod(b1,value)
	b2 := new(big.Int).Mul(c,pub_key)
        b2 = b2.Add(a2,b2)
        //b2 = b1.Mod(b2,value)
        b3 := new(big.Int).Mul(r,b2)
        //b3 = b1.Mod(b3,value)
        b4 := new(big.Int).Mul(c,pri_key2)
	b4 = b4.Add(a3,b4)
	//b4 = b4.Mod(b4,value)
	b5 := new(big.Int).Mul(r,b4)
	//b5 = b1.Mod(b5,value)     
	//fmt.Println("b1 : ",b1," b2 : ",b2," b3 : ",b3," b4 : ",b4," b5 : ",b5)
 	elapsed := time.Since(start)
        _,_ = f.WriteString("Signing took  " )
        f.WriteString(elapsed.String())	
        
 	verify_signature(f,mesg,b1,b2,b3,b4,b5,u1,u2,u3,u4,u5,u6,u7,u8,u9,v,T,tag,N)	        
}
