import {MyNumberNode} from '../src/node';

describe('MyNumberNode', () => {
  it('creates a new number node', () => {
    const node3 = new MyNumberNode(3);

    expect(node3.getValue()).toBe(3);
    expect(node3.getNext()).toBeNull;
  })
  it('links two nodes', () => {
    const node42 = new MyNumberNode(42);
    const node3 = new MyNumberNode(3, node42);

    expect(node3.getNext()).toBe(node42)
    expect(node42.getNext()).toBeNull
  })
  it('links three nodes', () => {
    const node42 = new MyNumberNode(42);
    const node21 = new MyNumberNode(21, node42);
    const node3 = new MyNumberNode(3, node21);

    expect(node3.getNext()).toBe(node21)
    const after3 = node3.getNext()
    if (after3 != null) {
      expect(after3.getNext()).toBe(node42)
    } else {
      expect(after3).not.toBeNull
    }
    
    expect(node42.getNext()).toBeNull
  })
  it('set new value', () => {
    const nodeA = new MyNumberNode(1);
    nodeA.setValue(42);
    expect(nodeA.getValue()).toBe(42)
  })
  it('reset link', () => {
    const node42 = new MyNumberNode(42);
    const node21 = new MyNumberNode(21, node42);
    const node3 = new MyNumberNode(3, node21);

    expect(node3.getNext()).toBe(node21)
    const after3 = node3.getNext()
    if (after3 != null) {
      expect(after3.getNext()).toBe(node42)
    } else {
      expect(after3).not.toBeNull
    }
    expect(node42.getNext()).toBeNull

    node3.setNext(node42);
    expect(node3.getNext()).toBe(node42)
  })
})