// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class HelloDB {

	static GONGSTRUCT_NAME = "Hello"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for pointers and slices of pointers declarations

	HelloPointersEncoding: HelloPointersEncoding = new HelloPointersEncoding
}

export class HelloPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}
