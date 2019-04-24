/*Anonymous endorsement of Transaction in Hyperledger Fabric : Fabric's Constant-sized Linkable Ring Signature*/
package main


import (
	"crypto/rand"
	"io"
	"math/big"
        "fmt"
        "os"
        "log"
	"time"       
)

type PrivateKey struct {
	P, Q    *big.Int
}

type PublicKey struct {
	N *big.Int
}

var base *big.Int
var bigOne = big.NewInt(1)
var bigTwo = big.NewInt(2)
var g1,h,t,y,s,z *big.Int
var start time.Time
var lambda = 1536
var msg_size int


func GenerateKey(random io.Reader) (*PublicKey, *PrivateKey, error) {
        li := lambda - 4
        mu := li - 2
        
	var x= new(big.Int).Exp(big.NewInt(2), big.NewInt(int64(li)), nil)
	var y= new(big.Int).Exp(big.NewInt(2), big.NewInt(int64(mu)),nil)
        
        for {

	     for {	
		p, err := rand.Prime(random,li)
                  
                
                   if err != nil {
			return nil, nil, err
		   }
                
                
		q, err := rand.Prime(random, li)
        	for {
		   if err != nil {
			return nil, nil, err
		   }
                   
                       

		//   fmt.Println(q.Cmp(new(big.Int).Sub(x,y)),q.Cmp(new(big.Int).Add(x,y)))
         	   if q.Cmp(new(big.Int).Sub(x,y))==1 && q.Cmp(new(big.Int).Add(x,y))==-1 {
			//  fmt.Println("Success")
                          break
                   }
                        
                   
                   //fmt.Println(p.Cmp(new(big.Int).Add(x,y)),p.Cmp(new(big.Int).Sub(x,y)))
                   q, err = rand.Prime(random,li)  
                }  
                	
		

                       privateKey := &PrivateKey{
				P:       p,
				Q:       q,
				
		       }
			N1 := new(big.Int).Mul(privateKey.P, privateKey.Q)
                        N1 = N1.Mul(bigTwo,N1)
                        N1 = N1.Add(bigOne,N1)
			//fmt.Println(N1.Cmp(new(big.Int).Sub(x1,y1)),N1.Cmp(new(big.Int).Add(x1,y1)))
			if N1.ProbablyPrime(2){
			    
                            publicKey := &PublicKey{
				N:       N1,     
			     
                            } 
                            return publicKey, privateKey, nil 		
                   	}  
			
		
		}	
                
              
         }
         

}







func main() {
          
        
        publicRSA, privateRSA, err := GenerateRSAInt(rand.Reader)
       // random io.Reader
        if err != nil {
		fmt.Printf("error")
	}
        fmt.Printf("RSA:%s,%s\n",publicRSA,privateRSA)
        //var N1 int64

        f, err := os.OpenFile("OUTPUT.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
        if err != nil {
        	log.Fatal(err)
    	}
    	


	fmt.Printf("Enter the endorsement set exponent size (set size be power of 2) : ")
    	//N1,err := strconv.ParseInt(os.Args[1], 10, 64)
        //if _, err := f.WriteString(os.Args[1]); err != nil {
        //	log.Fatal(err)
    	//}
         //_,_ = f.WriteString(" ")
        var N1 = big.NewInt(8)
	fmt.Println("Endorsers : ",N1)
	var N2 = new(big.Int).Exp(big.NewInt(2),N1,nil)
        N3 := N2.Uint64()
	var N int
        N = int(N3)
	publicKeyOriginaltest := make([]PublicKey,N)
	publicKeyOriginal := make([]PublicKey,N)
        privateKey := make([]PrivateKey,N)
	li:= lambda - 4

        for i:=0;i<N;i++ {   
		
        //	pk, sk, err2 := GenerateKey(rand.Reader)

	//	if err2 != nil {
	//		fmt.Printf("error")
	//	}
               		publicKeyOriginaltest[i].N,_=rand.Prime(rand.Reader,2*li) 
          //      privateKeytest[i].P=sk.P
           //     privateKeytest[i].Q=sk.Q
		
        	fmt.Println("Sno. : ",i+1," public key : ",publicKeyOriginaltest[i])
       }
      
      //var l = make([]big.Int,N)
      // base = GenerateBase(rand.Reader,publicRSA)
       //_,err = TransformPublicKey(rand.Reader,publicKeyTransformed,N4,new(big.Int).Div(publicRSA.N,big.NewInt(4)))
var counter int64
	
for counter=5 ; counter<=7 ; counter++ {
        
        N1 = big.NewInt(counter)
	//fmt.Println("Endorsers : ",N1)
	N2 = new(big.Int).Exp(big.NewInt(2),N1,nil)
        N3 := N2.Uint64()
	if _, err := f.WriteString("Endorser size: " + N2.String()); err != nil {
        	log.Fatal(err)
        }
        _,_ = f.WriteString(" ")
       
        N = int(N3)
	
        for i:=0;i<N;i++ {   
		
        	publicKeyOriginal[i].N=publicKeyOriginaltest[i].N
          //      privateKey[i].P=privateKeytest[i].P
            //    privateKey[i].Q=privateKeytest[i].Q
		
        	fmt.Println("Sno. : ",i+1," public key : ",publicKeyOriginaltest[i])
       }
     

fmt.Println("Enter the endorser id who wants to sign : ")
      var id int
      
      id1,_ := rand.Int(rand.Reader,new(big.Int).Sub(N2,bigOne))
      id = int(id1.Uint64())
      id = id + 1	

        pk,sk,_ := GenerateKey(rand.Reader)
	publicKeyOriginal[int(id)-1].N=pk.N
        privateKey[int(id)-1].P=sk.P
        privateKey[int(id)-1].Q=sk.Q
		//if err2 != nil {
		//	fmt.Printf("error")
		//}

     acc := new(big.Int).Set(base)
      
      for i := 0; i<N ;i++  {
	    	
			acc.Exp(acc,publicKeyOriginal[i].N,publicRSA.N)

		
	}
      
for msg_size=2048; msg_size<=8192 ; msg_size=msg_size*2 {
      Large := big.NewInt(int64(msg_size))
      _,_ = f.WriteString("Msg size :" + Large.String()) 
      for i:=1 ; i<=10 ; i++ {     
    
      
           message := make([]byte, msg_size)
           for i := 0; i < msg_size; i++ {
         
          message[i] = byte(65 + 20)  //A=65 and Z = 65+25
           }
      //return string(bytes)



       tid,_ := rand.Int(rand.Reader,big.NewInt(87476))
       mesg := string(message)
     
       //fmt.Println(mesg)
     
      fmt.Println("Witness Generation")
      w:= witness(publicKeyOriginal,N,int(id)-1,publicRSA.N)
      start = time.Now() 
      
             //key := new(big.Int).Mod(publicKeyOriginal[int(id)-1].N,value)  
        //w.Exp(w,publicKeyOriginal[int(id)-1].N,publicRSA.N)  
	//fmt.Println("acc : ",acc," w : ",w)
	//fmt.Println("g1 : ",g1," h : ",h," t : ",t," y : ",y," s : ",s," z : ",z)
        
	//fmt.Println("Signer constructs public values : ")
        txid := hash_txid(tid.String())
        g1 = new(big.Int).Exp(g1,txid,publicRSA.N)

       
        //fmt.Println("Tag gen")


        tag := GenTag(&privateKey[int(id)-1],publicRSA.N)
      	
	fmt.Println("Final Tag generated : ")
	

        //fmt.Println("Secret : ",privateKey[int(id)-1].P,"\n",privateKey[int(id)-1].Q)     
 	//fmt.Println("\nTag is : ",tag)
	//T := make([]*big.Int,5)
        T,r := GenPublicVal(rand.Reader,publicKeyOriginal[int(id)-1].N,&privateKey[int(id)-1],w,publicRSA.N)
/*	for i := range T {
		fmt.Println("T[] : ",T[i]) 	
	} */
	//fmt.Println(" r : ",r)
        proof_signer(f,mesg,rand.Reader,T,w,acc,r,publicKeyOriginal[int(id)-1].N,privateKey[int(id)-1].P,privateKey[int(id)-1].Q,tag,publicRSA.N)

    
   }
                  
   }
 }
 if err := f.Close(); err != nil {
	        log.Fatal(err)
	}
}



