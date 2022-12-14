import {MyNumberNode} from './node'
import {NodeType} from './node'

type initValueType = number | null

export class MyNumLinkedList {
  #head: NodeType

  constructor(value:initValueType = null) {
    if (value != null) {
      this.#head = new MyNumberNode(value)
    } else {
      this.#head = null
    }
  }

  getHead() {
    return this.#head
  }

  setNewHead(newValue: number) {
    const newNode = new MyNumberNode(newValue)
    if (this.#head != null) {
      newNode.setNext(this.#head)
    }
    this.#head = newNode
  }

  stringifyList(){
    let current: NodeType = this.getHead()
    let outputString = ""
    while (current != null) {
      outputString = outputString + current.getValue().toString()
      const nextNode = current.getNext()
      if (nextNode) {
        outputString = outputString + "-"
      }
      current = nextNode
    }
    return outputString
  }

  
  
}