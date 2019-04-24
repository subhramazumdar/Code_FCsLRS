/*Anonymous endorsement of Transaction in Hyperledger Fabric : Fabric's Constant-sized Linkable Ring Signature*/
package main

import (

	"math/big"
	"golang.org/x/crypto/sha3"
	"bytes"
        
)


func GenTag(privateKey *PrivateKey,N *big.Int) (*big.Int) {
	//  fmt.Println("Secret : ",privateKey.P,"\n",privateKey.Q)
          secret := new(big.Int).Add(privateKey.P,privateKey.Q)
          tag := new(big.Int).Exp(g1,secret,N)
	  return tag

}

func hash_txid(m string) (*big.Int) {
      var b bytes.Buffer
      b.WriteString(m)
      
// A hash needs to be 64 bytes long to have 256-bit collision resistance.
	hash := make([]byte, 16)
// Compute a 64-byte hash of buf and put it in h.
	sha3.ShakeSum128(hash, b.Bytes())	
	//fmt.Printf("%x\n", hash)
        buffer := bytes.NewBuffer(hash[:])
	c := new(big.Int).SetBytes(buffer.Bytes())
        return c
}
