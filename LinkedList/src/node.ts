export type NodeType = MyNumberNode | null

export class MyNumberNode {
  #value: number;
  #next: NodeType

  constructor(value:number, next:NodeType = null){
    this.#value = value;
    this.#next = next;
  }

  getValue(): number{
    return this.#value;
  }

  setValue(newValue:number) {
    this.#value = newValue;
  }

  getNext(): NodeType{
    return this.#next;
  }

  setNext(next: NodeType){
    this.#next = next;
  }
}