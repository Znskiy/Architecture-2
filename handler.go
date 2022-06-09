package lab2

import (
  "io"
  "bufio"
)

// ComputeHandler should be constructed with input io.Reader and output io.Writer.
// Its Compute() method should read the expression from input and write the computed result to the output.
type ComputeHandler struct {
  Input io.Reader
  Output io.Writer
}

func (ch *ComputeHandler) Compute() error {
  scanner := bufio.NewScanner(ch.Input)
  scanner.Scan()
  expression := scanner.Text()
  
  if scanner.Err() != nil {
    return scanner.Err()
  }

  resultExpression, calcErr := PrefixToPostfix(expression)
  if calcErr != nil {
    return calcErr
  }
  
  writer := bufio.NewWriter(ch.Output) 
  _, writeErr := writer.WriteString(resultExpression)
  if writeErr != nil {
    return writeErr
  }
  writer.Flush()
  return nil 
}
