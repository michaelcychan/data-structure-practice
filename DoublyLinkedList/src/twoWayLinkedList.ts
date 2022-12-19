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

  removeHead():(number | null) {
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
      if (nodeRemoval) {
        return nodeRemoval.getValue()
      } else {
        return null
      }
      
    }
    return null
  }

  removeTail():(number | null){
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
      if (nodeRemoval) {
        return nodeRemoval.getValue()
      } else {
        return null
      }
      
    }
    return null
  }

  removeByValue(value:number):(number | null) {
    let currentNode = this.#head;
    while (currentNode) {
      if (currentNode.getValue() == value) {
        break
      }
      currentNode = currentNode.getNext()
    }
    if (currentNode) {
      if (currentNode === this.#head) {
        this.removeHead()
      } else if (currentNode === this.#tail) {
        this.removeTail()
      } else {
        const prevNode = currentNode.getPrev()
        const nextNode = currentNode.getNext()

        if (prevNode) {
          prevNode.setNext(nextNode)
        }
        if (nextNode) {
          nextNode.setPrev(prevNode)
        }
      }
    }
    if (currentNode) {
      return currentNode.getValue()
    } else {
      return null
    }
    
  }
}