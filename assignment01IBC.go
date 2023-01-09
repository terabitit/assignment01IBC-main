//Package greetings shows the greetings
package assignment01IBC

import (
  "crypto/sha256"
  "fmt"
)

type BlockData struct {
Transactions []string
}
type Block struct {
Data        BlockData
PrevPointer *Block
PrevHash    string
CurrentHash string
}
func CalculateHash(inputBlock *Block) string {
  Transaction := fmt.Sprintf("%v",inputBlock.Data.Transactions)
  Result := fmt.Sprintf("%x\n", sha256.Sum256([]byte(Transaction)))
  return Result
}
func InsertBlock(dataToInsert BlockData, chainHead *Block) *Block {
  var newHead *Block = new(Block)
  if chainHead == nil{
    newHead.Data = dataToInsert
    newHead.CurrentHash = CalculateHash(newHead)

  }else{
    newHead.Data = dataToInsert
    newHead.CurrentHash = CalculateHash(newHead)
    newHead.PrevHash = chainHead.CurrentHash
    newHead.PrevPointer = chainHead
  }
  return newHead
}
func ChangeBlock(oldTrans string, newTrans string, chainHead *Block) {
  tempX := chainHead
  for tempX != nil{
      for i := 0 ; i < len(tempX.Data.Transactions) ; i++ {
        if tempX.Data.Transactions[i] == oldTrans{
          tempX.Data.Transactions[i] = newTrans
        }
      }
      tempX = tempX.PrevPointer
    }
  }
func ListBlocks(chainHead *Block) {
  tempX := chainHead
  for tempX != nil{
    fmt.Printf("%s\n",tempX.Data.Transactions)
    tempX = tempX.PrevPointer
  }
}
func VerifyChain(chainHead *Block) {

  tempX := chainHead
  for tempX != nil{
    if tempX.CurrentHash == CalculateHash(tempX){
    }else{
      fmt.Printf("Blockchain is compromised. Not verified")
      return
    }
    tempX = tempX.PrevPointer
  }
  fmt.Printf("Blockchain is not compromised. Blockchain verified")
}
