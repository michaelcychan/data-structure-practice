import {MyNumberNode} from './node'
import {NodeType} from './node'

type ValueType = number | null

export class MyNumLinkedList {
  #head: NodeType

  constructor(value:ValueType = null) {
    if (value != null) {
      this.#head = new MyNumberNode(value)
    } else {
      this.#head = null
    }
  }

  getHead() {
    return this.#head
  }

  addNewHead(newValue: number) {
    const newNode = new MyNumberNode(newValue)
    if (this.#head != null) {
      newNode.setNext(this.#head)
    }
    this.#head = newNode
  }

  stringifyList(){
    let current: NodeType = this.getHead()
    let outputString = ""
    while (current) {
      outputString += current.getValue().toString()
      const nextNode = current.getNext()
      if (nextNode) {
        outputString += "-"
      }
      current = nextNode
    }
    return outputString
  }

  setNewHead(newNode: NodeType) {
    this.#head = newNode
  }

  removeNode(value:number) {
    let current: NodeType = this.getHead()
    if (current) {
      if (current.getValue() === value) {
        this.setNewHead(current.getNext())
        return
      }
    }
    while (current) {
      const nextNode = current.getNext()

      if (nextNode) {
        if (nextNode.getValue() === value) {
          current.setNext(nextNode.getNext())
          break
        } else {
          current = current.getNext()
        }
      } else {
        console.error(`target: ${value} not found, nothing was removed`)
        break
      }
    }
  }

  reverseIndex(fromLast:number):ValueType {
    let tailPointer = this.getHead();
    let targetPointer = this.getHead();
    const head = this.getHead()
    let counter = 0;
    while (tailPointer) {
      tailPointer = tailPointer.getNext()
      counter += 1
      if (fromLast < counter ) {
        if (targetPointer) {
          targetPointer = targetPointer.getNext()
        }
      }
    }
    if (counter == fromLast && head) {
      return head.getValue()
    } else if (counter < fromLast) {
      return null
    }
    if (targetPointer) {
      console.log(targetPointer.getValue())
      return targetPointer.getValue()
    } else {
      return null
    }
  }
  
  
}