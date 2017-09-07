package resolved

func encodedSize(encoding string) int {
	switch encoding {
	case "float":
		return 4
	case "uint16":
		return 2
	case "uint8":
		return 1
	default:
		panic(encoding)
	}
}

func sizeInBytes(t *ScalarType) int {
	switch t.Name {
	case "float":
		return 4
	default:
		panic(t)
	}
}

func align(offset int, size int) int {
	return (offset + size - 1) & ^(size - 1)
}

func calcOffsets(file *File) {
	for _, v := range file.Vertex {
		firstByteSize := 1
		offset := 0
		for i, comp := range v.Components {
			encodedScalarByteSize := encodedSize(comp.Encoding)
			switch t := comp.Type.(type) {
			case *VectorType:
				scalarByteSize := encodedScalarByteSize
				if i == 1 {
					firstByteSize = scalarByteSize
				}
				offset = align(offset, scalarByteSize)
				comp.ByteOffset = offset
				comp.ByteSize = scalarByteSize * t.Width

				offset += comp.ByteSize
			default:
				panic(t)
			}
		}
		v.ByteSize = align(offset, firstByteSize)
	}
}

func splitDataAndSamplers(fields []*Field) ([]*Field, []*Field) {
	uniformData := []*Field{}
	uniformSamplers := []*Field{}
	for _, f := range fields {
		if isSampler(f.Type) {
			uniformSamplers = append(uniformSamplers, f)
		} else {
			uniformData = append(uniformData, f)
		}
	}
	return uniformData, uniformSamplers
}
