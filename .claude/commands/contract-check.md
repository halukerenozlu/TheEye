---
description: Event sözleşmesindeki sapmayı AGENTS.md, Go servisleri ve dashboard tipleri arasında kontrol et.
---

`Event` sözleşmesinin tanımları arasında sapma olup olmadığını kontrol et. Yalnızca oku; istenmedikçe hiçbir şeyi düzenleme.

## Kaynaklar (AGENTS.md 19. maddedeki öncelik sırasıyla)

1. `AGENTS.md` 5. madde — zorunlu ve önerilen alanlar (sözleşme)
2. `docs/API.md` — belgelenmiş yanıt şekli
3. `services/collector/models/event.go` — ingestion modeli
4. `services/api/cmd/api/main.go` — API'nin fiilen serialize ettiği
5. `apps/dashboard/src/types/index.ts` — frontend'in inandığı

## Rapor

Her alan için bir satır, her kaynak için bir sütun içeren tek bir tablo üret. Sonra ayrı ayrı listele:

- **Bir kaynakta eksik olan zorunlu alanlar** — AGENTS.md 5. madde ihlali
- **Sözleşmede olmayıp kaynakta olan alanlar** — belgelenmemiş eklemeler
- **Aynı mantıksal alan için ad veya tip uyuşmazlıkları**
- **`severity` vs `severity_level`**: hangi katman hangisini kullanıyor, ayrım kasıtlı mı

## Kurallar

- Sözleşme koddan üstündür. Kod ile `AGENTS.md` çelişiyorsa bu, güncellenecek bir spec değil, raporlanacak bir sapmadır.
- İstenmedikçe düzeltme önerme. Farkı raporla.
- Yığın zaten ayaktaysa canlı yanıtla doğrula:
  `curl -s "http://localhost:8080/v1/events?limit=1"`
  Sırf bunun için Docker başlatma; kapalıysa öyle söyle.
- AGENTS.md 16. maddedeki verdict ile bitir: Accept / Accept with minimal patch / Rework needed / Reject.
