package macroman

import (
	"bytes"
	"testing"
)

// string containing all UTF-8 characters
const utf8_test = "DGƒ˙˝ä‚gãì¥È^≠∞≤Úaû[s°ÀY´‹6œ”\"§ˇ.<`∏π∫kÆ’òRzØÌ⁄)_mÖO¨»ﬁÏ!?Bõ£]Ù Kñ–p3V|≥ÂÎÒù‘Ë#0CWà5@dföÃÊv'2UoÇêßAwyÔÛ˛NÉîª„>çèº‰x≈Œ}®Ω&F/PXÜó™æø›\\ôﬂuåï©ErÕˆ*-¢µ«7Tc~â1†¶÷¯+¬…HJn‡Í$%,Mıé8:tú˘(LSQlü∆;q¿ÿÁj€4Zí¡Ó =∑9h{—◊Ÿei˚IÑ±“bÄÅá∂˜¸ë•√·"

// string containing all MacRoman characters as octal values.
// Obviously, macroman_test[n] corresponds to utf8_test[n]  
const macroman_test = "\104\107\304\372\375\212\342\147\213\223\264\351\136\255\260\262\362\141\236\133\163\241\313\131\253\334\066\317\323\042\244\377\056\074\140\270\271\272\153\256\325\230\122\172\257\355\332\051\137\155\205\117\254\310\336\354\041\077\102\233\243\135\364\040\113\226\320\360\160\063\126\174\263\345\353\361\235\324\350\043\060\103\127\210\065\100\144\146\232\314\346\166\047\062\125\157\202\220\247\101\167\171\357\363\376\116\203\224\273\343\076\215\217\274\344\170\305\316\175\250\275\046\106\057\120\130\206\227\252\276\277\335\134\231\337\165\214\225\251\105\162\315\366\052\055\242\265\307\067\124\143\176\211\061\240\246\326\370\053\302\311\110\112\156\340\352\044\045\054\115\365\216\070\072\164\234\371\050\114\123\121\154\237\306\073\161\300\330\347\152\333\064\132\222\301\356\312\075\267\071\150\173\321\327\331\145\151\373\111\204\261\322\142\200\201\207\266\367\374\221\245\303\341"


func TestEncode(t *testing.T) {
	dst := make([]byte, len(utf8_test))
	l := Encode(dst, []byte(utf8_test))
	if l != len(macroman_test) {
		t.Errorf("Encode() length mismatch, expected %d, got %d", len(macroman_test), l)
	}
	if string(dst[:l]) != macroman_test {
		t.Errorf("Encode() mismatch, expected %s, got %s", macroman_test, string(dst))
	}
}

func TestDecode(t *testing.T) {
	dst := make([]byte, len(utf8_test))
	l, r, err := Decode(dst, []byte(macroman_test))
	if err != nil {
		t.Error(err)
	}
	if r != len(macroman_test) {
		t.Errorf("Decode() read length mismatch, expected %d, got %d", len(macroman_test), l)
	}
	if l != len(utf8_test) {
		t.Errorf("Decode() length mismatch, expected %d, got %d", len(utf8_test), l)
	}
	if string(dst[:l]) != utf8_test {
		t.Errorf("Decode() mismatch, expected %s, got %s", utf8_test, string(dst))
	}
}

func TestEncoder(t *testing.T) {
	buf := bytes.NewBuffer(make([]byte, 0))
	w := NewEncoder(buf)
	l, err := w.Write([]byte(utf8_test))
	if err != nil {
		t.Error(err)
	}	
	if l != len(macroman_test) {
		t.Errorf("encoder.Write() length mismatch, expected %d, got %d", len(macroman_test), l)
	}
	dst := buf.Bytes()
	if string(dst[:l]) != macroman_test {
		t.Errorf("encoder.Write() mismatch, expected %s, got %s", macroman_test, string(dst))
	}
}

func TestDecoder(t *testing.T) {
	buf := bytes.NewBuffer([]byte(macroman_test))
	dst := make([]byte, len(utf8_test))
	r := NewDecoder(buf)
 	l, err := r.Read(dst)
	if err != nil {
		t.Error(err)
	}	
	if l != len(utf8_test) {
		t.Errorf("decoder.Read() length mismatch, expected %d, got %d", len(utf8_test), l)
	}
	if string(dst[:l]) != utf8_test {
		t.Errorf("decoder.Read() mismatch, expected %s, got %s", utf8_test, string(dst))
	}
}