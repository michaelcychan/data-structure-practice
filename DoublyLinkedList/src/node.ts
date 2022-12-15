type ValueType = number | null
type NodeType = MyTwoWayNode | null

export class MyTwoWayNode {
  #value: ValueType
  #nextNode: NodeType
  #prevNode: NodeType
  
  constructor(value: ValueType = null) {
    this.#value = value
    this.#nextNode = null
    this.#prevNode = null
  }

  getValue(){
    return this.#value
  }

  getNext(){
    return this.#nextNode
  }

  getPrev(){
    return this.#prevNode
  }

  setNext(nextNode: NodeType){
    this.#nextNode = nextNode
  }

  setPrev(prevNode: NodeType){
    this.#prevNode = prevNode
  }
}