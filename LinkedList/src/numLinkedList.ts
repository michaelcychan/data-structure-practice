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

  stringifyList(){

  }

  
  
}