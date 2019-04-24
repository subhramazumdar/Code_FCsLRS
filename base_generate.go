/*Anonymous endorsement of Transaction in Hyperledger Fabric : Fabric's Constant-sized Linkable Ring Signature*/
package main


import (
	"crypto/rand"
	"io"
	"math/big"
        "fmt"
       
)

func GenerateBase(random io.Reader,publicRSA *PublicRSA) (*big.Int){
	                   for {
				base, _ = rand.Int(random, new(big.Int).Div(publicRSA.N,big.NewInt(4)))
				fmt.Println("Generator")
				g := new(big.Int).GCD(nil, nil, base,new(big.Int).Div(publicRSA.N,big.NewInt(4)))
		              	if g.Cmp(bigOne) == 0 {
					return base	
				}
			   }
}

