package keygen

func GenerateKey(flags *KeyGeneratorFlags) string {
	return ProcessKeyString(flags.Ks)
}