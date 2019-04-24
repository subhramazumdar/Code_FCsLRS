/*Anonymous endorsement of Transaction in Hyperledger Fabric : Fabric's Constant-sized Linkable Ring Signature*/
package main

import (
	
        "crypto/rand"
	"io"
	"math/big"
)
func GenPublicVal(random io.Reader,secret *big.Int,privateKey *PrivateKey,witness *big.Int, N *big.Int ) ([]*big.Int,*big.Int){

       var T []*big.Int
       T=make([]*big.Int,5)	
       value := new(big.Int).Div(N,big.NewInt(4))
       value = new(big.Int).Sub(value,big.NewInt(1))	
       r,_ := rand.Int(random,value) 	
       T[0]=new(big.Int).Exp(g1,r,N)
       T[1]=new(big.Int).Exp(h,r,N)
       temp := new(big.Int).Add(secret,r)
       temp =new(big.Int).Exp(z,temp,N)
       T[1]=new(big.Int).Mul(T[1],temp)
       T[1]=new(big.Int).Mod(T[1],N)

       T[2]=new(big.Int).Exp(s,r,N)
       temp=new(big.Int).Exp(g1,privateKey.Q,N)
       T[2]=new(big.Int).Mul(T[2],temp)
       T[2]=new(big.Int).Mod(T[2],N)
       temp=new(big.Int).Exp(y,r,N)
       T[3]=new(big.Int).Mul(witness,temp)
       T[3]=new(big.Int).Mod(T[3],N)
       
       temp1:=new(big.Int).Mul(big.NewInt(2),privateKey.P)
       temp=new(big.Int).Exp(g1,temp1,N)
       
       T[4]=new(big.Int).Exp(t,r,N)
       T[4]=new(big.Int).Mul(T[4],temp)
       T[4]=new(big.Int).Mod(T[4],N)
       return T,r	

}
