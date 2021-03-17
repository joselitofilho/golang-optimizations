# Cenários de Memory Leaking

Normalmente, ao programar em uma linguagem que oferece suporte à coleta de lixo automática, não precisamos nos preocupar com problemas de vazamento de memória, pois a limpeza regular da memória não utilizada será realizada. No entanto, devemos ter em mente alguns cenários especiais que podem causar vazamento de memória. A seguir, veremos alguns cenários.

Para entender sobre a gestão de memória do Go e o coletor de lixo veja esse [link](https://deepu.tech/memory-management-in-golang/).

## Causado por substrings

O compilador/runtime padrão do Go permite que uma string `S` e uma substring de `S` compartilhem o mesmo bloco de memória subjacente. Isso é muito bom para economizar memória e processamento. Mas às vezes pode causar vazamento de memória.

Por exemplo, depois que a função `KindofLeakingCausedBySubstring` no exemplo a seguir é chamada, haverá cerca de 1MB de memória vazando, até que a variável de nível de pacote `str0` seja modificada novamente em outro lugar.

```Go
var str0 string

func KindofLeakingCausedBySubstring() {
	str := allocMemory(1 << 20) // 1MB
	fn(str)
}

func allocMemory(size int) string {
	return string(make([]byte, size))
}

func fn(str1 string) {
	str0 = str1[:50]
}
```
