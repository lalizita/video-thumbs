# Passo a passo para construir a aplicação

- [ ] Selecionar um vídeo .mp4
- [ ] Criar função que gera a playlist HLS
- [ ] Criar função que gera imagens a partir da HLS
- Servir estáticos com echo
- Criar Go routines


```
curl -v http://localhost:8000/thumbs/segment_260.png
curl -v localhost:8000/coelho/playlist.m3u8
```