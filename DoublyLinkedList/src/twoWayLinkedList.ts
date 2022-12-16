import {NodeType, MyTwoWayNode} from './node'

export class TwoWayLinkedList {

  #head: NodeType;
  #tail: NodeType;

  constructor(value:number | null = null) {
    if (value) {
      const firstNode: NodeType = new MyTwoWayNode(value)
      this.#head = firstNode;
      this.#tail = firstNode;
    } else {
      this.#head = null
      this.#tail = null
    }
  }

  getHead() {
    return this.#head
  }

  getTail() {
    return this.#tail
  }

  stringify():string{
    let output:string = "";
    let node = this.#head;
    while (node) {
      const value = node.getValue()
      if (value) {
        output += value.toString()
      }
      if (node.getNext()) {
        output += "-"
      }
      node = node.getNext()
    }
    return output
  }

  addHead(newValue: number) {
    const newNode = new MyTwoWayNode(newValue)
    newNode.setNext(this.#head)
    if (this.#head) {
      this.#head.setPrev(newNode)
    }
    this.#head = newNode
    
    if (!this.#tail) {
      this.#tail = newNode
    }

  }

  addTail(value:number) {
    const newTail = new MyTwoWayNode(value);
    newTail.setPrev(this.#tail)
    if (this.#tail) {
      this.#tail.setNext(newTail)
    }
    this.#tail = newTail
    if (!this.#head) {
      this.#head = newTail
    }
  }

  length():number {
    let counter:number = 0;
    let node = this.#head;
    while (node) {
      counter += 1;
      node = node.getNext()
    }
    return counter;
  }

  removeHead():NodeType {
    if (this.#head) {
      const nodeRemoval = this.#head
      const newHead = this.#head.getNext()
      this.#head = newHead
      if (newHead) {
        newHead.setPrev(null)
      }
      if (nodeRemoval === this.#tail) {
        this.#tail = null
      }
      return nodeRemoval
    }
    return null
  }

  removeTail():NodeType{
    if (this.#tail) {
      const nodeRemoval = this.#tail
      const newTail = this.#tail.getPrev()
      this.#tail = newTail
      if (newTail) {
        newTail.setNext(null)
      }
      if (nodeRemoval === this.#head) {
        this.#head = null
      }
      return nodeRemoval
    }
    return null
  }

  removeByValue() {

  }
}