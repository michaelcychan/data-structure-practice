export type NodeType = MyTwoWayNode | null

export class MyTwoWayNode {
  #value: number
  #nextNode: NodeType
  #prevNode: NodeType
  
  constructor(value: number) {
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