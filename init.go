/*Anonymous endorsement of Transaction in Hyperledger Fabric : Fabric's Constant-sized Linkable Ring Signature*/
package main


import (
	"crypto/rand"
	"io"
	"fmt"
	"math/big"
       

       
)

type PrivateRSA struct {
	P, Q    *big.Int
        Totient *big.Int
}

type PublicRSA struct {
	N *big.Int
}

func GenerateRSAInt(random io.Reader) (*PublicRSA, *PrivateRSA, error) {
    	
        //var q,err1 
        
	for {
                
		p, err := rand.Prime(random, lambda)
                for {			
		if err != nil {
			return nil, nil, err
		}
		p = new(big.Int).Mul(p, big.NewInt(2))
		p = new(big.Int).Add(p, big.NewInt(1))
		//fmt.Println(p)
		// Probability of p being non-prime is negligible
               
		if p.ProbablyPrime(10) {
		     break
                }
                p, err = rand.Prime(random, lambda)
	       }
                q, err1 := rand.Prime(random, lambda)
        	for {	
		if err1 != nil {
		        return nil, nil, err1
		}
                q = new(big.Int).Mul(q, big.NewInt(2))
		q = new(big.Int).Add(q, big.NewInt(1))
		//fmt.Println(p)
		// Probability of p being non-prime is negligible
               
		 if q.ProbablyPrime(10) {
                     break
	         }
                 q, err = rand.Prime(random, lambda)
                }
	        pminus1 := new(big.Int).Sub(p, bigOne)
		qminus1 := new(big.Int).Sub(q, bigOne)
		totient := new(big.Int).Mul(pminus1, qminus1) 
			privateRSA := &PrivateRSA{
				P:       p,
				Q:       q,
				Totient : totient,
		
			}
			N1 := new(big.Int).Mul(privateRSA.P, privateRSA.Q)
                        
                          
			publicRSA := &PublicRSA{
				N:       N1,     
				 
                        }
			for {        		
				a, err1 := rand.Int(random, N1)
				fmt.Println("Generator")
                                fmt.Println(a)
				if err1 != nil {
					return nil,nil, err1
				}
                                x := new(big.Int).Exp(a,big.NewInt(2),N1)
				// Ensure generator order is not 2 (efficiency)
				if x.Cmp(big.NewInt(1)) == 0 {
					continue
				}
				
					
				// Return if generator order is q
				g := new(big.Int).GCD(nil, nil,new(big.Int).Sub(x,big.NewInt(1)),N1)
		              if g.Cmp(bigOne) == 0 {
				
					base = x
					g1 = new(big.Int).Exp(base,big.NewInt(31),N1)
					h = new(big.Int).Exp(base,big.NewInt(5),N1)
					t = new(big.Int).Exp(base,big.NewInt(23),N1)
					y = new(big.Int).Exp(base,big.NewInt(13),N1)
					s = new(big.Int).Exp(base,big.NewInt(47),N1)
					z = new(big.Int).Exp(base,big.NewInt(71),N1)	
					break

			     }
			}
			return publicRSA, privateRSA, nil

		      
        		
	}
}





