/*Anonymous endorsement of Transaction in Hyperledger Fabric : Fabric's Constant-sized Linkable Ring Signature*/
package main


import (
	
	"math/big"
        
       
)



func witness(publicKeyOriginal []PublicKey,N int, id int,publicRSA *big.Int) (*big.Int) {
        acc := new(big.Int).Set(base)
       // value := new(big.Int).Div(publicRSA,big.NewInt(4)) 
	for i :=0 ; i<N ;i++  {
	       if i==id {
			continue
		}
		acc.Exp(acc,publicKeyOriginal[i].N,publicRSA)
	}
       
	//w := new(big.Int).Mod(secret,value)
        //val := new(big.Int).Mul()
       
        //fmt.Println("acc : ",acc)
	return acc
}




 
